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
	  "_id": "653acea3531031bd1d4a1cd9",
	  "key": "avatar/user/1698352803936",
	  "hash": "FlmF1mwgVC8DGSQKwgMm0Jj2Fpgc",
	  "fsize": 27336,
	  "bucket": "cow-image",
	  "name": "d75533b7-743f-11ee-8638-00155d48097e"
	}
	*/
	cos := NewCos(config)
	bucket := cos.NewVideoBucket()
	uid := utils.GetUUID()
	fn := utils.GetGenerateUniqueCode()
	f, _ := os.OpenFile("/own/go_place/InvertedCow/ui/src/assets/avatar/avatar.png", os.O_RDONLY, 0777)
	err := bucket.PutFile(path.Join(UserAvatarPath, fn), uid, f)
	assert.Nil(t, err)
}
