package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type PostController interface {
	Post(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (p *postController) Post(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	originText := ctx.PostForm("originText")
	user, ok := ctx.Get("user")
	if !ok {
		result.Error(e.ErrBadRequest)
		return
	}
	userId := user.(*dto.UserInfo).ID
	var hasSource bool
	hasSourceRaw := strings.ToLower(ctx.PostForm("hasSource"))
	if hasSourceRaw == "true" {
		hasSource = true
	}
	token, err := p.postService.Post(ctx, originText, int64(userId), hasSource)
	if err != nil {
		// TODO: record and report
		result.Error(e.ErrBadRequest)
		return
	}
	result.SuccessData(token)
}
