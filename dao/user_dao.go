package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type UserDao interface {
	// InsertUser 创建用户
	InsertUser(db *gorm.DB, user *po.User) error
}

type userDao struct {
}

func NewUserDao() UserDao {
	return &userDao{}
}

func (s *userDao) InsertUser(db *gorm.DB, user *po.User) error {
	return db.Create(user).Error
}
