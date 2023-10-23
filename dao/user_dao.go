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
	// GetUserByEmail 根据用户邮箱获取用户信息
	GetUserByEmail(db *gorm.DB, email string) (*po.User, error)
	// GetUserByLoginName 根据用户登录名称获取用户
	GetUserByLoginName(db *gorm.DB, loginName string) (*po.User, error)
}

type userDao struct {
}

func NewUserDao() UserDao {
	return &userDao{}
}

func (u *userDao) InsertUser(db *gorm.DB, user *po.User) error {
	return db.Create(user).Error
}

func (u *userDao) UpdateUser(db *gorm.DB, user *po.User) error {
	return db.Model(user).Updates(user).Error
}

func (u *userDao) DeleteUserByID(db *gorm.DB, id uint) error {
	return db.Delete(&po.User{}, id).Error
}

func (u *userDao) GetUserByID(db *gorm.DB, id uint) (*po.User, error) {
	user := &po.User{}
	err := db.First(&user, id).Error
	return user, err
}

func (u *userDao) GetUserByEmail(db *gorm.DB, email string) (*po.User, error) {
	user := &po.User{}
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userDao) GetUserByLoginName(db *gorm.DB, loginName string) (*po.User, error) {
	user := &po.User{}
	err := db.Where("login_name = ?", loginName).First(&user).Error
	return user, err
}
