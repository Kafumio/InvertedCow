package service

import (
	"InvertedCow/dao"
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type AccountService interface {
	SignUp(user *po.User) error
}

type accountService struct {
	db      *gorm.DB
	userDao dao.UserDao
}

func NewAccountService(db *gorm.DB, userDao dao.UserDao) AccountService {
	return &accountService{
		db:      db,
		userDao: userDao,
	}
}

func (a *accountService) SignUp(user *po.User) error {
	return nil
}
