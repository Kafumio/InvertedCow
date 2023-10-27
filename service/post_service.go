package service

import (
	conf "InvertedCow/config"
	"InvertedCow/data"
	"context"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type SourceService interface {
	Upload(ctx context.Context) error // 回调，主要是绑定业务属性
}

type sourceService struct {
	config *conf.AppConfig
	db     *gorm.DB
	cos    *data.Cos
	redis  *redis.Client
}

func NewSourceService(config *conf.AppConfig,
	db *gorm.DB, cos *data.Cos, redis *redis.Client) SourceService {
	return &sourceService{
		config: config,
		db:     db,
		cos:    cos,
		redis:  redis,
	}
}

// Upload TODO: 回调
func (s *sourceService) Upload(ctx context.Context) error {
	return nil
}
