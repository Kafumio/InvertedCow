package config

import "gopkg.in/ini.v1"

type CosConfig struct {
	AccessKey   string `ini:"accessKey"`
	SecretKey   string `ini:"secretKey"`
	Region      string `ini:"region"`
	ImageBucket string `ini:"imageBucket"`
}

func NewCosConfig(cfg *ini.File) *CosConfig {
	cosConfig := &CosConfig{}
	cfg.Section("email").MapTo(cosConfig)
	return cosConfig
}
