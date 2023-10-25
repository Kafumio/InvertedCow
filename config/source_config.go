package config

import "gopkg.in/ini.v1"

type SourceConfig struct {
	AccessKey string `ini:"access_key"`
	SecretKey string `ini:"secret_key"`
	Bucket    string `ini:"bucket"`
}

func newSourceConfig(conf *ini.File) *SourceConfig {
	sourceConfig := &SourceConfig{}
	conf.Section("source").MapTo(sourceConfig)
	return sourceConfig
}
