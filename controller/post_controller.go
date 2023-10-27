package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type PostController interface {
	Post(ctx *gin.Context)
	Upload(ctx *gin.Context)
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
	token, err := p.postService.Post(ctx, originText, userId, hasSource)
	if err != nil {
		// TODO: record and report
		result.Error(e.ErrBadRequest)
		return
	}
	result.SuccessData(token)
}

// Upload 提供给七牛云的回调，用于绑定业务属性
func (p *postController) Upload(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	source := &dto.Source{}
	err := ctx.ShouldBindJSON(source)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	uid, err := strconv.Atoi(source.Name)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	err = p.postService.Upload(ctx, source.Id, source.Hash, source.Key, source.Bucket, uint(uid), source.FSize)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	result.SuccessMessage("ok") // TODO
}
