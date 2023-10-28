package data

import (
	"InvertedCow/config"
	"github.com/go-redis/redis"
)

func NewRedisClient(conf *config.AppConfig) *redis.Client {
	cfg := conf.RedisConfig
	redis := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0, // 数据库
	})
	return redis
}
