package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"InvertedCow/utils"
	"github.com/gin-gonic/gin"
)

// AccountController 关于一些账号信息的handler
type AccountController interface {
	// SendAuthCode 发送验证码
	SendAuthCode(ctx *gin.Context)
	// SignIn 用户登录
	SignIn(ctx *gin.Context)
	// SignUp 用户注册
	SignUp(ctx *gin.Context)
	// GetUserInfo 从token里面读取用户信息
	GetUserInfo(ctx *gin.Context)
}

type accountController struct {
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

func (a *accountController) SendAuthCode(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	email := ctx.PostForm("email")
	kind := ctx.PostForm("type")
	if email != "" && !utils.VerifyEmailFormat(email) {
		result.SimpleErrorMessage("邮箱格式错误")
		return
	}
	// 生成code
	_, err := a.accountService.SendAuthCode(email, kind)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("验证码发送成功")
}

func (a *accountController) SignUp(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	user := &po.User{
		Email:    ctx.PostForm("email"),
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
	}
	code := ctx.PostForm("code")
	// check username
	if len(user.Username) < 3 {
		result.SimpleErrorMessage("用户名过短")
		return
	}
	// check password
	if len(user.Password) < 5 {
		result.SimpleErrorMessage("用户密码过短")
		return
	}
	err := a.accountService.SignUp(user, code)
	if err != nil {
		result.Error(err)
	} else {
		result.SuccessMessage("注册成功")
	}
}

func (a *accountController) SignIn(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	//获取并检验用户参数
	kind := ctx.PostForm("type")
	account := ctx.PostForm("account")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	code := ctx.PostForm("code")
	if kind != "password" && kind != "email" {
		result.Error(e.ErrBadRequest)
		return
	} else if kind == "password" && (account == "" || password == "") {
		result.Error(e.ErrBadRequest)
		return
	} else if kind == "email" && (email == "" || code == "") {
		result.Error(e.ErrBadRequest)
		return
	}
	// 登录
	var token string
	var err *e.Error
	if kind == "password" {
		token, err = a.accountService.PasswordSignIn(account, password)
	} else if kind == "email" {
		token, err = a.accountService.EmailSignIn(email, code)
	} else {
		result.Error(e.ErrSignInType)
		return
	}
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(token)
}

func (u *accountController) GetUserInfo(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	user := ctx.Keys["user"].(*dto.UserInfo)
	result.SuccessData(user)
}
