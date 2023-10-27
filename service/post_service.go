package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
	"InvertedCow/data"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type PostService interface {
	Post(ctx context.Context, text string, publisher uint, onlyText bool) (*dto.Token, error)
	Upload(ctx context.Context, id, hash, key, bucket string, pid uint, fSize int64) error // 回调，主要是绑定业务属性
	Deprecated()
}

var (
	once sync.Once
)

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
	post := &postService{
		config: config,
		db:     db,
		cos:    cos,
		redis:  redis,
		pd:     pd,
		sd:     sd,
	}
	once.Do(func() {
		go post.Deprecated()
	})
	return post
}

func (p *postService) Post(ctx context.Context, text string, userId uint, onlyText bool) (*dto.Token, error) {
	var err error
	post := &po.Post{
		State:     1,
		Publisher: userId,
		Text:      text,
	}
	// 单纯文字，目前不存在这种情况，只是便于测试
	if onlyText {
		post.State = 2
		err = p.pd.InsertPost(p.db, post)
		if err != nil {
			return nil, err
		}
		return &dto.Token{
			Token: "",
			PID:   post.ID,
		}, nil
	}
	// 生成授权Token
	err = p.pd.InsertPost(p.db, post)
	if err != nil {
		return nil, err
	}

	bucket := p.cos.NewVideoBucket()

	token := &dto.Token{
		Token: bucket.Token(),
		PID:   post.ID,
	}
	return token, nil
}

// Upload
// docs
// 1. 存储source info
// 2. 关联业务属性
// 3. 返回响应
// TODO: transaction
func (p *postService) Upload(ctx context.Context, id, hash, key, bucket string, pid uint, fSize int64) error {
	source := &po.Source{
		PostID:   pid, // origin_post_uid
		FileName: id,
		Hash:     hash,
		Size:     fSize,
		Key:      key,
		Bucket:   bucket,
	}
	err := p.sd.InsertSource(p.db, source)
	if err != nil {
		return err
	}
	post, err := p.pd.GetPostByID(p.db, pid)
	if err != nil {
		return err
	}
	post.State = 2 // TODO: 目前只支持单source发布
	err = p.pd.UpdatePost(p.db, post)
	if err != nil {
		return err
	}
	return nil
}

func (p *postService) Deprecated() {
	ticker := time.NewTimer(2 * time.Second)
	for {
		<-ticker.C
		fmt.Println("ggg")
		list, err := p.pd.GetPostListWithoutPage(p.db, &po.Post{
			State: 1,
		})
		if err != nil {
			return
		}
		pIds := make([]uint, len(list))
		for i := 0; i < len(list); i++ {
			pIds[i] = list[i].ID
		}
		sources, err := p.sd.GetSourcesByPostIds(p.db, pIds)
		if err != nil {
			log.Println("deprecated error", err)
			return
		}
		successIds := make(map[uint]struct{})
		for i := 0; i < len(sources); i++ {
			successIds[sources[i].PostID] = struct{}{}
			post, err := p.pd.GetPostByID(p.db, sources[i].PostID)
			if err != nil {
				return
			}
			post.State = 2
			err = p.pd.UpdatePost(p.db, post)
			if err != nil {
				return
			}
		}
		for i := 0; i < len(list); i++ {
			if _, ok := successIds[list[i].ID]; !ok {
				// 更改为失败
				post, err := p.pd.GetPostByID(p.db, sources[i].PostID)
				if err != nil {
					return
				}
				post.State = 3
				err = p.pd.UpdatePost(p.db, post)
				if err != nil {
					return
				}
			}
		}
	}
}
