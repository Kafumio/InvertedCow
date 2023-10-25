package service

import (
	conf "InvertedCow/config"
	"InvertedCow/model/dto"
	"context"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type SourceService interface {
	Token(ctx context.Context) (dto.Token, error) // 获取Token，同时记录 唯一业务标识
	Upload(ctx context.Context) error             // 回调，主要是绑定业务属性
}

type sourceService struct {
	config *conf.AppConfig
	db     *gorm.DB
	redis  *redis.Client
}

func NewSourceService(config *conf.AppConfig,
	db *gorm.DB, redis *redis.Client) SourceService {
	return &sourceService{
		config: config,
		db:     db,
		redis:  redis,
	}
}

func (s *sourceService) Token(ctx context.Context) (dto.Token, error) {
	return dto.Token{}, nil
}

func (s *sourceService) Upload(ctx context.Context) error {
	return nil
}
