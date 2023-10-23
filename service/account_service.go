package service

import (
	"InvertedCow/dao"
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/utils"
	"gorm.io/gorm"
)

type AccountService interface {
	SignUp(user *po.User) error
	PasswordSignIn(account string, password string) (string, *e.Error)
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

func (a *accountService) PasswordSignIn(account string, password string) (string, *e.Error) {
	var user *po.User
	var userErr error
	if utils.VerifyEmailFormat(account) {
		user, userErr = a.userDao.GetUserByEmail(a.db, account)
	} else {
		user, userErr = a.userDao.GetUserByLoginName(a.db, account)
	}
	// 没有该用户
	if user == nil || userErr == gorm.ErrRecordNotFound {
		return "", e.ErrUserNotExist
	}
	if userErr != nil {
		return "", e.ErrMysql
	}
	// 比较密码
	if !utils.ComparePwd(user.Password, password+user.Salt) {
		return "", e.ErrUserNameOrPasswordWrong
	}
	userInfo := dto.NewUserInfo(user)
	token, err := utils.GenerateToken(utils.Claims{
		ID:        userInfo.ID,
		Avatar:    userInfo.Avatar,
		Username:  userInfo.Username,
		LoginName: userInfo.LoginName,
		Email:     userInfo.Email,
	})
	if err != nil {
		return "", e.ErrUserUnknownError
	}
	return token, nil
}
