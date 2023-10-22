package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
)

type Controller struct {
	AccountController AccountController
}

func NewController(
	accountController AccountController,
) *Controller {
	return &Controller{
		AccountController: accountController,
	}
}
