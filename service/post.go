package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
	"InvertedCow/data"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/utils"
	"context"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type PostService interface {
	Post(ctx context.Context, originText string, publisher int64, hasSource bool) (*dto.Token, error)
}

type postService struct {
	config *conf.AppConfig
	db     *gorm.DB
	cos    *data.Cos
	redis  *redis.Client
	pd     dao.PostDao
}

func NewPostService(config *conf.AppConfig,
	db *gorm.DB, cos *data.Cos, redis *redis.Client, pd dao.PostDao) PostService {
	return &postService{
		config: config,
		db:     db,
		cos:    cos,
		redis:  redis,
		pd:     pd,
	}
}

func (p *postService) Post(ctx context.Context, originText string, userId int64, hasSource bool) (*dto.Token, error) {
	var err error
	uid := utils.GetUUID()
	post := &po.Post{
		State:     1,
		Publisher: userId,
		Text:      originText,
	}
	if !hasSource {
		post.State = 2
		err = p.pd.InsertPost(p.db, post)
		if err != nil {
			return nil, err
		}
		return &dto.Token{
			Token:     "",
			OriginUrl: uid,
		}, nil
	}
	// 生成授权Token
	// TODO: 时限监听 —— 连接 Token时效，避免上传成功但动态发布失败。
	bucket := p.cos.NewVideoBucket()

	token := &dto.Token{
		Token:     bucket.Token(),
		OriginUrl: uid,
	}
	err = p.pd.InsertPost(p.db, post)
	if err != nil {
		return nil, err
	}
	return token, nil
}
