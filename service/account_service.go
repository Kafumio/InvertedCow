package service

import (
	conf "InvertedCow/config"
	"InvertedCow/dao"
	"InvertedCow/data"
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/utils"
	"errors"
	"github.com/Chain-Zhang/pinyin"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"mime/multipart"
	"path"
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
	// UploadAvatar 上传头像
	UploadAvatar(file *multipart.FileHeader) (string, *e.Error)
	// GetAccountInfo 读取账号信息
	GetAccountInfo(ctx *gin.Context) (*dto.AccountInfo, *e.Error)
	// ChangePassword 修改用户密码
	ChangePassword(ctx *gin.Context, oldPassword string, newPassword string) *e.Error
	// UpdateAccount 更新账号信息
	UpdateAccount(ctx *gin.Context, user *po.User) *e.Error
}

type accountService struct {
	config  *conf.AppConfig
	db      *gorm.DB
	redis   *redis.Client
	cos     *data.Cos
	userDao dao.UserDao
}

func NewAccountService(config *conf.AppConfig,
	db *gorm.DB, redis *redis.Client, cos *data.Cos, userDao dao.UserDao) AccountService {
	return &accountService{
		config:  config,
		db:      db,
		redis:   redis,
		cos:     cos,
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
	// 设置出生日期默认值
	t := time.Time{}
	if user.BirthDay == t {
		user.BirthDay = time.Now()
	}
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
		ID: user.ID,
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
		ID: user.ID,
	})
	if err != nil {
		return "", e.ErrUserUnknownError
	}
	return token, nil
}

func (a *accountService) ChangePassword(ctx *gin.Context, oldPassword string, newPassword string) *e.Error {
	userInfo := ctx.Keys["user"].(*dto.UserInfo)
	//检验用户名
	user, err := a.userDao.GetUserByID(a.db, userInfo.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return e.ErrUserNotExist
	}
	if err != nil {
		return e.ErrMysql
	}
	//检验旧密码
	if !utils.ComparePwd(user.Password, oldPassword+user.Salt) {
		return e.ErrUserNameOrPasswordWrong
	}
	password, getPwdErr := utils.GetPwd(newPassword + user.Salt)
	if getPwdErr != nil {
		return e.ErrPasswordEncodeFailed
	}
	user.Password = string(password)
	err = a.userDao.UpdateUser(a.db, user)
	if err != nil {
		return e.ErrMysql
	}
	return nil
}

func (a *accountService) UpdateAccount(ctx *gin.Context, user *po.User) *e.Error {
	userInfo := ctx.Keys["user"].(*dto.UserInfo)
	user.ID = userInfo.ID
	// 不能更新账号名称和密码
	user.LoginName = ""
	user.Password = ""
	err := a.userDao.UpdateUser(a.db, user)
	if err != nil {
		return e.ErrMysql
	}
	return nil
}

func (a *accountService) GetAccountInfo(ctx *gin.Context) (*dto.AccountInfo, *e.Error) {
	user := ctx.Keys["user"].(*dto.UserInfo)
	userInfo, err := a.userDao.GetUserByID(a.db, user.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.ErrUserNotExist
	}
	return dto.NewAccountInfo(userInfo), nil
}

const (
	// UserAvatarPath cos中，用户图片存储的位置
	UserAvatarPath = "/avatar/user"
)

func (a *accountService) UploadAvatar(file *multipart.FileHeader) (string, *e.Error) {
	bucket := a.cos.NewImageBucket()
	fileName := file.Filename
	fileName = utils.GetUUID() + "." + path.Base(fileName)
	file2, err := file.Open()
	if err != nil {
		return "", e.ErrBadRequest
	}
	err = bucket.PutFileSimple(path.Join(UserAvatarPath, fileName), file2)
	if err != nil {
		return "", e.ErrServer
	}
	return bucket.MakeUrl(a.config.ImageProUrl, path.Join(UserAvatarPath, fileName)), nil
}
