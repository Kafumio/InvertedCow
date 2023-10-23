package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
	e "InvertedCow/error"
	"InvertedCow/model/po"
	"InvertedCow/utils"
	"github.com/Chain-Zhang/pinyin"
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
	SignUp(user *po.User, code string) *e.Error
	// PasswordSignIn 用户密码登录
	PasswordSignIn(account string, password string) (string, *e.Error)
	// EmailSignIn 邮件登录
	EmailSignIn(email string, code string) (string, *e.Error)
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
		redis:   redis,
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

func (a *accountService) SignUp(user *po.User, code string) *e.Error {
	// 检测是否已注册过
	f, err := a.userDao.CheckEmail(a.db, user.Email)
	if f {
		return e.ErrUserEmailIsExist
	}
	// 检测code
	result := a.redis.Get(SignUpEmailProKey + user.Email)
	if result.Err() != nil {
		return e.ErrUserUnknownError
	}
	if result.Val() != code {
		return e.ErrSignUpCodeWrong
	}
	// 生成用户名称，唯一
	loginName, err := pinyin.New(user.Username).Split("").Convert()
	if err != nil {
		return e.ErrUserUnknownError
	}
	loginName = loginName + utils.GetRandomNumber(3)
	for i := 0; i < 5; i++ {
		b, err := a.userDao.CheckLoginName(a.db, user.LoginName)
		if err != nil {
			return e.ErrMysql
		}
		if b {
			loginName = loginName + utils.GetRandomNumber(1)
		} else {
			break
		}
	}
	user.LoginName = loginName
	//进行注册操作
	salt := utils.GetRandomPassword(32)
	user.Salt = salt
	newPassword, err := utils.GetPwd(user.Password + salt)
	if err != nil {
		return e.ErrPasswordEncodeFailed
	}
	user.Password = string(newPassword)

	err = a.userDao.InsertUser(a.db, user)
	if err != nil {
		return e.ErrMysql
	}
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
	token, err := utils.GenerateToken(utils.Claims{
		ID:        user.ID,
		Avatar:    user.Avatar,
		Username:  user.Username,
		LoginName: user.LoginName,
		Email:     user.Email,
	})
	if err != nil {
		return "", e.ErrUserUnknownError
	}
	return token, nil
}

func (a *accountService) EmailSignIn(email string, code string) (string, *e.Error) {
	if !utils.VerifyEmailFormat(email) {
		return "", e.ErrUserEmailIsNotValid
	}
	// 获取用户
	user, err := a.userDao.GetUserByEmail(a.db, email)
	if err != nil {
		return "", e.ErrMysql
	}
	if err == gorm.ErrRecordNotFound {
		return "", e.ErrUserNotExist
	}
	// 检测验证码
	key := SignInEmailProKey + email
	result, err2 := a.redis.Get(key).Result()
	if err2 != nil {
		return "", e.ErrUserUnknownError
	}
	if result != code {
		return "", e.ErrSignInCodeWrong
	}
	token, err := utils.GenerateToken(utils.Claims{
		ID:        user.ID,
		Avatar:    user.Avatar,
		Username:  user.Username,
		LoginName: user.LoginName,
		Email:     user.Email,
	})
	if err != nil {
		return "", e.ErrUserUnknownError
	}
	return token, nil
}
