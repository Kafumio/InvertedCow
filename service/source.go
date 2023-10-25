package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
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
	config  *conf.AppConfig
	db      *gorm.DB
	redis   *redis.Client
	userDao dao.SourceDao
}
