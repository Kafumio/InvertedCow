package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type PostDao interface {
	InsertPost(db *gorm.DB, post *po.Post) error
}

type postDao struct {
}

func NewPostDao() PostDao {
	return &postDao{}
}

func (p *postDao) InsertPost(db *gorm.DB, post *po.Post) error {
	return db.Create(post).Error
}
