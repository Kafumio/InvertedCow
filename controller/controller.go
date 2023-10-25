package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
	NewSourceController,
)

type Controller struct {
	AccountController AccountController
	SourceController  SourceController
}

func NewController(
	accountController AccountController,
	SourceController SourceController,
) *Controller {
	return &Controller{
		AccountController: accountController,
		SourceController:  SourceController,
	}
}
