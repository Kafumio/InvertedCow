package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
	NewPostController,
	NewViewController,
)

type Controller struct {
	AccountController AccountController
	PostController    PostController
}

func NewController(
	accountController AccountController,
	postController PostController,
) *Controller {
	return &Controller{
		AccountController: accountController,
		PostController:    postController,
	}
}
