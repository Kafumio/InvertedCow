package service

import (
	"InvertedCow/config"
	"InvertedCow/dao"
	"InvertedCow/data"
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
)

type ViewService interface {
	// NextPost 通过算法读取下一个推荐视频的id
	NextPost(ctx *gin.Context) (uint, *e.Error)
	// PrePost 返回上一个视频 todo:是否由前端缓存实现？
	PrePost(ctx *gin.Context) (uint, *e.Error)
	// GetPostById 读取视频信息
	GetPostById(postId uint) (*dto.PostDtoForView, *e.Error)
}

type viewService struct {
	db        *gorm.DB
	cos       *data.Cos
	config    *config.AppConfig
	postDao   dao.PostDao
	sourceDao dao.SourceDao
	userDao   dao.UserDao
}

func NewViewService(db *gorm.DB, cos *data.Cos,
	ac *config.AppConfig, pd dao.PostDao, sd dao.SourceDao, ud dao.UserDao) ViewService {
	return &viewService{
		db:        db,
		cos:       cos,
		config:    ac,
		postDao:   pd,
		sourceDao: sd,
		userDao:   ud,
	}
}

// todo：实现推荐算法推荐用户视频
func (v *viewService) NextPost(ctx *gin.Context) (uint, *e.Error) {
	count, err := v.postDao.GetPostCount(v.db, &po.Post{
		State: 2,
	})
	if err != nil {
		return 0, e.ErrMysql
	}
	r := rand.Intn(int(count))
	posts, err := v.postDao.GetPostList(v.db, &dto.PageQuery{
		Query:    &po.Post{State: 2},
		Page:     r,
		PageSize: 1,
	})
	if err != nil || len(posts) == 0 {
		return 0, e.ErrMysql
	}
	return posts[0].ID, nil
}

// todo：实现推荐算法推荐用户视频
func (v *viewService) PrePost(ctx *gin.Context) (uint, *e.Error) {
	count, err := v.postDao.GetPostCount(v.db, &po.Post{
		State: 2,
	})
	if err != nil {
		return 0, e.ErrMysql
	}
	r := rand.Intn(int(count))
	posts, err := v.postDao.GetPostList(v.db, &dto.PageQuery{
		Query:    &po.Post{State: 2},
		Page:     r,
		PageSize: 1,
	})
	if err != nil || len(posts) == 0 {
		return 0, e.ErrMysql
	}
	return posts[0].ID, nil
}

func (v *viewService) GetPostById(postId uint) (*dto.PostDtoForView, *e.Error) {
	post, err := v.postDao.GetPostByID(v.db, postId)
	if err == gorm.ErrRecordNotFound {
		return nil, e.ErrMysql
	}
	if err != nil {
		return nil, e.ErrMysql
	}
	sourceDto := dto.NewPostDtoForView(post)

	// 读取视频资源
	source, err := v.sourceDao.GetSourceByPostId(v.db, post.ID)
	if err != nil {
		return nil, e.ErrMysql
	}
	bucket := v.cos.NewVideoBucket()
	sourceDto.VideoUrl = bucket.MakeUrl(v.config.VideoProUrl, source.Key)
	// 读取用户信息
	user, err := v.userDao.GetUserByID(v.db, post.Publisher)
	sourceDto.PublisherAvatar = user.Avatar
	return sourceDto, nil
}
