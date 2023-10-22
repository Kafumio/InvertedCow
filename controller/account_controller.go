package controller

import (
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
)

// AccountController 关于一些账号信息的handler
type AccountController interface {
	SignUp(ctx *gin.Context)
}

type accountController struct {
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &accountController{
		accountService: accountService,
	}
}

func (a *accountController) SignUp(ctx *gin.Context) {

}
