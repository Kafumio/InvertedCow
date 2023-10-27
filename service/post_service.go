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
	Upload(ctx context.Context, id, hash, key, bucket, uid string, fSize int64) error // 回调，主要是绑定业务属性
}

type postService struct {
	config *conf.AppConfig
	db     *gorm.DB
	cos    *data.Cos
	redis  *redis.Client
	pd     dao.PostDao
	sd     dao.SourceDao
}

func NewPostService(config *conf.AppConfig,
	db *gorm.DB, cos *data.Cos, redis *redis.Client, pd dao.PostDao, sd dao.SourceDao) PostService {
	return &postService{
		config: config,
		db:     db,
		cos:    cos,
		redis:  redis,
		pd:     pd,
		sd:     sd,
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
	// TODO: 时限监听 —— 连接 Token 时效，避免上传成功但动态发布失败。
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

// Upload
// docs
// 1. 存储source info
// 2. 关联业务属性
// 3. 返回响应
// TODO: transaction
func (p *postService) Upload(ctx context.Context, id, hash, key, bucket, uid string, fSize int64) error {
	source := &po.Source{
		UID:      uid, // origin_post_uid
		FileName: id,
		Hash:     hash,
		Size:     fSize,
		Key:      key, // TODO: source type, recording to the suffix of the key. Or get from request
		Bucket:   bucket,
	}
	err := p.sd.InsertSource(p.db, source)
	if err != nil {
		// TODO: log
		return err
	}
	post, err := p.pd.GetPostByUID(p.db, uid)
	if err != nil {
		return err
	}
	post.State = 2 // TODO: 目前只支持单source发布。
	err = p.pd.UpdatePost(p.db, post)
	if err != nil {
		return err
	}
	return nil
}
