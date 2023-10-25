package config

import "gopkg.in/ini.v1"

type CosConfig struct {
	AccessKey   string `ini:"accessKey"`
	SecretKey   string `ini:"secretKey"`
	Region      string `ini:"region"`
	ImageBucket string `ini:"imageBucket"`
	ImageProUrl string `ini:"imageProUrl"`
}

func NewCosConfig(cfg *ini.File) *CosConfig {
	cosConfig := &CosConfig{}
	cfg.Section("cos").MapTo(cosConfig)
	return cosConfig
}
