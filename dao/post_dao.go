package dao

import (
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type PostDao interface {
	InsertPost(db *gorm.DB, post *po.Post) error
	UpdatePost(db *gorm.DB, post *po.Post) error
	// GetPostCount 读取post的数量
	GetPostCount(db *gorm.DB, post *po.Post) (int64, error)
	// GetPostList 读取post列表
	GetPostList(db *gorm.DB, pageQuery *dto.PageQuery) ([]*po.Post, error)
	GetPostListWithoutPage(db *gorm.DB, post *po.Post) ([]*po.Post, error)
	// GetPostByID 根据ID获取指定动态
	GetPostByID(db *gorm.DB, postId uint) (*po.Post, error)
}

type postDao struct {
}

func NewPostDao() PostDao {
	return &postDao{}
}

func (p *postDao) InsertPost(db *gorm.DB, post *po.Post) error {
	return db.Create(post).Error
}

func (p *postDao) UpdatePost(db *gorm.DB, post *po.Post) error {
	return db.Model(post).Updates(post).Error
}

func (p *postDao) GetPostCount(db *gorm.DB, post *po.Post) (int64, error) {
	var count int64
	if post != nil && post.Text != "" {
		db = db.Where("text LIKE ?", "%"+post.Text+"%")
	}
	if post != nil && post.State != 0 {
		db = db.Where("state = ?", post.State)
	}
	if post != nil && post.Publisher != 0 {
		db = db.Where("publisher LIKE ?", post.Publisher)
	}
	err := db.Model(&po.Post{}).Count(&count).Error
	return count, err
}

func (p *postDao) GetPostList(db *gorm.DB, pageQuery *dto.PageQuery) ([]*po.Post, error) {
	var post *po.Post
	if pageQuery.Query != nil {
		post = pageQuery.Query.(*po.Post)
	}
	offset := (pageQuery.Page - 1) * pageQuery.PageSize
	var posts []*po.Post
	if post != nil && post.Text != "" {
		db = db.Where("text LIKE ?", "%"+post.Text+"%")
	}
	if post != nil && post.State != 0 {
		db = db.Where("state = ?", post.State)
	}
	if post != nil && post.Publisher != 0 {
		db = db.Where("publisher = ?", post.Publisher)
	}
	db = db.Limit(pageQuery.PageSize).Offset(offset)
	if pageQuery.SortProperty != "" && pageQuery.SortRule != "" {
		order := pageQuery.SortProperty + " " + pageQuery.SortRule
		db = db.Order(order)
	}
	err := db.Find(&posts).Error
	return posts, err
}

func (p *postDao) GetPostByID(db *gorm.DB, postId uint) (*po.Post, error) {
	var post po.Post
	err := db.Find(&post, postId).Error
	return &post, err
}

func (p *postDao) GetPostListWithoutPage(db *gorm.DB, post *po.Post) ([]*po.Post, error) {
	var posts []*po.Post
	if post.State != 0 {
		db = db.Where(`state = ?`, post.State)
	}
	db = db.Where(`created_at <= DATE_SUB(NOW(), INTERVAL 1 HOUR)`)
	err := db.Find(&posts).Error
	return posts, err
}
