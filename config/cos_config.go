package config

import "gopkg.in/ini.v1"

type CosConfig struct {
	AccessKey   string `ini:"accessKey"`
	SecretKey   string `ini:"secretKey"`
	Region      string `ini:"region"`
	ImageBucket string `ini:"imageBucket"`
	ImageProUrl string `ini:"imageProUrl"`
	VideoBucket string `ini:"videoBucket"`
	VideoProUrl string `ini:"videoProUrl"`
}

func newCosConfig(cfg *ini.File) *CosConfig {
	cosConfig := &CosConfig{}
	cfg.Section("cos").MapTo(cosConfig)
	return cosConfig
}
