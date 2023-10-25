package data

import (
	conf "InvertedCow/config"
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"os"
	"strings"
)

type Cos struct {
	cosConfig *conf.CosConfig
	mac       *qbox.Mac
}

func NewCos(config *conf.AppConfig) *Cos {
	cosConfig := config.CosConfig
	return &Cos{
		cosConfig: config.CosConfig,
		mac:       qbox.NewMac(cosConfig.AccessKey, cosConfig.SecretKey),
	}
}

func (c *Cos) NewImageBucket() *Bucket {
	return c.NewBucket(c.cosConfig.ImageBucket)
}

func (c *Cos) NewBucket(bucketName string) *Bucket {
	return &Bucket{
		cosConfig: c.cosConfig,
		mac:       c.mac,
		putPolicy: storage.PutPolicy{
			Scope: bucketName,
		},
	}
}

type Bucket struct {
	cosConfig *conf.CosConfig
	mac       *qbox.Mac
	putPolicy storage.PutPolicy
}

// PutFile 上传文件，以数据流的形式
func (b *Bucket) PutFile(key string, reader io.Reader) {
	// 生成上传屏障
	upToken := b.putPolicy.UploadToken(b.mac)
	cfg := storage.Config{}
	// 空间对应的机房
	if b.cosConfig.Region == "ZoneHuadong" {
		cfg.Region = &storage.ZoneHuadong
	}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}
	total, err := GetReaderLen(reader)
	if err != nil {
		return
	}
	err = formUploader.Put(context.Background(), &ret, upToken, key, reader.(io.ReaderAt), total, &putExtra)
	if err != nil {
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

func GetReaderLen(reader io.Reader) (length int64, err error) {
	switch v := reader.(type) {
	case *bytes.Buffer:
		length = int64(v.Len())
	case *bytes.Reader:
		length = int64(v.Len())
	case *strings.Reader:
		length = int64(v.Len())
	case *os.File:
		stat, ferr := v.Stat()
		if ferr != nil {
			err = fmt.Errorf("can't get reader length: %s", ferr.Error())
		} else {
			length = stat.Size()
		}
	case *io.LimitedReader:
		length = int64(v.N)
	default:
		err = fmt.Errorf("can't get reader content length, unkown reader type")
	}
	return
}
