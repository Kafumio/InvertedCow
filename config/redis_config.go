package config

import "gopkg.in/ini.v1"

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
}

func newRedisConfig(cfg *ini.File) *RedisConfig {
	redisConfig := &RedisConfig{}
	cfg.Section("redis").MapTo(redisConfig)
	return redisConfig
}
