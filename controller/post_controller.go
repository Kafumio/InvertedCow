package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"sync"
)

type PostController interface {
	Post(ctx *gin.Context)
	Upload(ctx *gin.Context)
}

var (
	once sync.Once
)

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	once.Do(func() {
		postService.Deprecated()
	})
	return &postController{
		postService: postService,
	}
}

func (p *postController) Post(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	text := ctx.PostForm("text")
	user, ok := ctx.Get("user")
	if !ok {
		result.Error(e.ErrBadRequest)
		return
	}
	userId := user.(*dto.UserInfo).ID
	var onlyText bool
	onlyTextRaw := strings.ToLower(ctx.PostForm("onlyText"))
	if onlyTextRaw == "true" {
		onlyText = true
	}
	token, err := p.postService.Post(ctx, text, userId, onlyText)
	if err != nil {
		log.Println("Post service error", err)
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
	// 数字转字符串
	pid, err := strconv.Atoi(source.PID)
	if err != nil {
		log.Println("Atoi error", err)
		result.Error(e.ErrBadRequest)
		return
	}
	err = p.postService.Upload(ctx, source.Id, source.Hash, source.Key, source.Bucket, uint(pid), source.FSize)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	result.SuccessMessage("ok") // TODO
}
