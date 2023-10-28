package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
	NewPostController,
	NewViewController,
	NewRelationController,
)

type Controller struct {
	AccountController  AccountController
	PostController     PostController
	ViewController     ViewController
	RelationController RelationController
}

func NewController(
	accountController AccountController,
	postController PostController,
	viewController ViewController,
	relationController RelationController,
) *Controller {
	return &Controller{
		AccountController:  accountController,
		PostController:     postController,
		ViewController:     viewController,
		RelationController: relationController,
	}
}
