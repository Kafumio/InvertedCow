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

func (c *Cos) NewBucket(bucket, returnBody, CallbackURL, CallbackBody, CallbackBodyType string) *Bucket {
	return &Bucket{
		cosConfig: c.cosConfig,
		mac:       c.mac,
		putPolicy: storage.PutPolicy{
			Scope:            bucket,
			ReturnBody:       returnBody,
			CallbackURL:      CallbackURL,
			CallbackBody:     CallbackBody,
			CallbackBodyType: CallbackBodyType,
		},
	}
}

func (c *Cos) NewBucketOnlyBucket(bucket string) *Bucket {
	return c.NewBucket(bucket, "", "", "", "")
}

func (c *Cos) NewImageBucket() *Bucket {
	return c.NewBucketOnlyBucket(c.cosConfig.ImageBucket)
}

func (c *Cos) NewVideoBucket() *Bucket {
	return c.NewBucket(c.cosConfig.VideoBucket, "", c.cosConfig.UploadCallback,
		// name 为 uid，唯一标识一条动态，需要客户端传入
		`{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		"application/json")
}

type Bucket struct {
	cosConfig *conf.CosConfig
	proUrl    string
	mac       *qbox.Mac
	putPolicy storage.PutPolicy
}

// PutFileSimple 上传文件，以数据流的形式
func (b *Bucket) PutFileSimple(key string, reader io.Reader) error {
	return b.PutFile(key, "", reader)
}

func (b *Bucket) PutFile(key, uid string, reader io.Reader) error {
	// 生成上传屏障
	upToken := b.putPolicy.UploadToken(b.mac)
	cfg := storage.Config{}
	// 空间对应的机房
	if b.cosConfig.Region == "ZoneHuanan" {
		cfg.Region = &storage.ZoneHuanan
	}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	var putExtra = storage.RputV2Extra{}
	if len(uid) > 0 {
		putExtra.CustomVars = map[string]string{
			"x:name": uid,
		}
	}
	total, err := GetReaderLen(reader)
	if err != nil {
		return err
	}
	// 判断key是否以/开头
	if len(key) > 0 && key[0] == '/' {
		key = key[1:]
	}
	return formUploader.Put(context.Background(), &ret, upToken, key, reader.(io.ReaderAt), total, &putExtra)
}

// MakeUrl 生成访问对象存储的url
// key 文件路径
// proUrl
func (b *Bucket) MakeUrl(proUrl string, key string) string {
	// 判断key是否以/开头
	if len(key) > 0 && key[0] == '/' {
		key = key[1:]
	}
	publicAccessURL := storage.MakePublicURL(proUrl, key)
	return publicAccessURL
}

// GetReaderLen 读取reader的长度
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
	case *LimitedReadCloser:
		length = int64(v.N)
	case FixedLengthReader:
		length = v.Size()
	default:
		err = fmt.Errorf("can't get reader content length, unkown reader type")
	}
	return
}

func (b *Bucket) Token() string {

	return b.putPolicy.UploadToken(b.mac)
}

type FixedLengthReader interface {
	io.Reader
	Size() int64
}

type LimitedReadCloser struct {
	io.LimitedReader
}
