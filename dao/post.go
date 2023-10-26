package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type PostDao interface {
	InsertPost(db *gorm.DB, post *po.Post) error
	GetPostByUID(db *gorm.DB, uid string) (*po.Post, error)
	UpdatePost(db *gorm.DB, post *po.Post) error
}

type postDao struct {
}

func NewPostDao() PostDao {
	return &postDao{}
}

func (p *postDao) InsertPost(db *gorm.DB, post *po.Post) error {
	return db.Create(post).Error
}

func (p *postDao) GetPostByUID(db *gorm.DB, uid string) (*po.Post, error) {
	var post po.Post
	err := db.Where("origin_url = ?", uid).First(&post).Error
	return &post, err
}

func (p *postDao) UpdatePost(db *gorm.DB, post *po.Post) error {
	return db.Model(post).Updates(post).Error
}
