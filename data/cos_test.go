package data

import (
	config2 "InvertedCow/config"
	"InvertedCow/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

const (
	UserAvatarPath = "/avatar/user"
)

func TestUploadCallback(t *testing.T) {
	config := &config2.AppConfig{
		CosConfig: &config2.CosConfig{
			// fill this gap.
		},
	}
	/**
	{
	  "_id": "653b772c531031bd1d4a1cdb",
	  "key": "avatar/user/4cb4b6e7-74a4-11ee-adfb-00155d48097e", // 文件路径
	  "hash": "FlmF1mwgVC8DGSQKwgMm0Jj2Fpgc",
	  "fsize": 27336,
	  "bucket": "cow-image",
	  "pid": "1698395950996" // 动态ID
	}
	*/
	cos := NewCos(config)
	bucket := cos.NewVideoBucket()
	fn := utils.GetUUID()
	pid := utils.GetGenerateUniqueCode() // 数字转字符串 动态ID
	f, _ := os.OpenFile("/own/go_place/InvertedCow/ui/src/assets/avatar/avatar.png", os.O_RDONLY, 0777)
	err := bucket.PutFile(path.Join(UserAvatarPath, fn), pid, f)
	assert.Nil(t, err)
}
