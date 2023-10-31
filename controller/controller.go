package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewController,
	NewAccountController,
	NewPostController,
	NewViewController,
	NewRelationController,
	NewCommentController,
)

type Controller struct {
	AccountController  AccountController
	PostController     PostController
	ViewController     ViewController
	RelationController RelationController
	CommentController  CommentController
}

func NewController(
	accountController AccountController,
	postController PostController,
	viewController ViewController,
	relationController RelationController,
	commentController CommentController,
) *Controller {
	return &Controller{
		AccountController:  accountController,
		PostController:     postController,
		ViewController:     viewController,
		RelationController: relationController,
		CommentController:  commentController,
	}
}
