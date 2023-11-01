package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ViewController interface {
	// NextPost 通过算法读取下一个推荐视频的id
	NextPost(ctx *gin.Context)
	// PrePost 返回上一个视频 todo:是否由前端缓存实现？
	PrePost(ctx *gin.Context)
	// GetPostById 读取视频信息
	GetPostById(ctx *gin.Context)
	// LikePost 给动态点赞
	LikePost(ctx *gin.Context)
}

type viewController struct {
	viewService service.ViewService
}

func NewViewController(vs service.ViewService) ViewController {
	return &viewController{
		viewService: vs,
	}
}

func (v *viewController) NextPost(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	postId, err := v.viewService.NextPost(ctx)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(postId)

}

func (v *viewController) PrePost(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	postId, err := v.viewService.PrePost(ctx)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(postId)
}

func (v *viewController) GetPostById(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	postIdStr := ctx.Param("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	post, err2 := v.viewService.GetPostById(ctx, uint(postId))
	if err2 != nil {
		result.Error(err2)
		return
	}
	result.SuccessData(post)
}

func (v *viewController) LikePost(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	postIdStr := ctx.Param("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	err2 := v.viewService.LikePost(ctx, uint(postId))
	if err2 != nil {
		result.Error(err2)
		return
	}
	result.SuccessMessage("请求成功")
}
