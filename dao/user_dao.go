package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type UserDao interface {
	// InsertUser 创建用户
	InsertUser(db *gorm.DB, user *po.User) error
	// UpdateUser 更新用户，不更新零值
	UpdateUser(db *gorm.DB, user *po.User) error
	// DeleteUserByID 删除用户
	DeleteUserByID(db *gorm.DB, id uint) error
	// GetUserByID 通过用户id获取用户, 找不到会返回异常而不是nil
	GetUserByID(db *gorm.DB, id uint) (*po.User, error)
}

type userDao struct {
}

func NewUserDao() UserDao {
	return &userDao{}
}

func (s *userDao) InsertUser(db *gorm.DB, user *po.User) error {
	return db.Create(user).Error
}

func (s *userDao) UpdateUser(db *gorm.DB, user *po.User) error {
	return db.Model(user).Updates(user).Error
}

func (s *userDao) DeleteUserByID(db *gorm.DB, id uint) error {
	return db.Delete(&po.User{}, id).Error
}

func (s *userDao) GetUserByID(db *gorm.DB, id uint) (*po.User, error) {
	user := &po.User{}
	err := db.First(&user, id).Error
	return user, err
}
