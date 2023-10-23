package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/utils"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"time"
)

const (
	// SignUpEmailProKey  注册时email的前缀key
	SignUpEmailProKey = "emailcode-signUp-"
	// SignInEmailProKey 登录时email的前缀key
	SignInEmailProKey = "emailcode-signIn-"
)

type AccountService interface {
	// SendAuthCode 获取邮件的验证码
	SendAuthCode(email string, kind string) (string, *e.Error)
	// SignUp 用户注册
	SignUp(user *po.User) error
	// PasswordSignIn 用户登录
	PasswordSignIn(account string, password string) (string, *e.Error)
}

type accountService struct {
	config  conf.AppConfig
	db      *gorm.DB
	redis   *redis.Client
	userDao dao.UserDao
}

func NewAccountService(config conf.AppConfig,
	db *gorm.DB, redis *redis.Client, userDao dao.UserDao) AccountService {
	return &accountService{
		config:  config,
		db:      db,
		userDao: userDao,
	}
}

func (a *accountService) SendAuthCode(email string, kind string) (string, *e.Error) {
	if kind == "signUp" {
		f, err := a.userDao.CheckEmail(a.db, email)
		if err != nil {
			return "", e.ErrUserUnknownError
		}
		if f {
			return "", e.ErrUserEmailIsExist
		}
	}

	var subject string
	if kind == "signUp" {
		subject = "InvertedCow注册验证码"
	} else if kind == "signIn" {
		subject = "InvertedCow登录验证码"
	}
	// 发送code
	code := utils.GetRandomNumber(6)
	message := utils.EmailMessage{
		To:      []string{email},
		Subject: subject,
		Body:    "验证码：" + code,
	}
	err := utils.SendMail(a.config.EmailConfig, message)
	if err != nil {
		return "", e.ErrUserUnknownError
	}
	// 存储到redis
	var key string
	if kind == "signUp" {
		key = SignUpEmailProKey + email
	} else {
		key = SignInEmailProKey + email
	}
	_, err2 := a.redis.Set(key, code, 10*time.Minute).Result()
	if err2 != nil {
		return "", e.ErrUserUnknownError
	}
	return code, nil
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
