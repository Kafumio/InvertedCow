package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

// PostDao
// TODO: crud
type PostDao interface {
	InsertPost(db *gorm.DB, post *po.Post) error
	GetPostByID(db *gorm.DB, pid uint) (*po.Post, error)
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

func (p *postDao) GetPostByID(db *gorm.DB, pid uint) (*po.Post, error) {
	var post po.Post
	err := db.First(&post, pid).Error
	return &post, err
}

func (p *postDao) UpdatePost(db *gorm.DB, post *po.Post) error {
	return db.Model(post).Updates(post).Error
}
