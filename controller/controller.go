package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
	NewRelationController,
)

type Controller struct {
	AccountController  AccountController
	RelationController RelationController
}

func NewController(
	accountController AccountController,
	relationController RelationController,
) *Controller {
	return &Controller{
		AccountController:  accountController,
		RelationController: relationController,
	}
}
