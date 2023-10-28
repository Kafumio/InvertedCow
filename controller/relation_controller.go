package controller

import (
	"InvertedCow/model/dto"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RelationController 关于用户之间关系的一些handler
type RelationController interface {
	// AddFollow 用户关注
	AddFollow(ctx *gin.Context)
	// CancelFollow 用户取关
	CancelFollow(ctx *gin.Context)
	// GetFollowList 获取关注列表
	GetFollowList(ctx *gin.Context)
	// GetFollowerList 获取粉丝列表
	GetFollowerList(ctx *gin.Context)
}

type relationController struct {
	relationService service.RelationService
}

func NewRelationController(relationService service.RelationService) RelationController {
	return &relationController{
		relationService: relationService,
	}
}

func (r *relationController) AddFollow(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.Keys["user"].(*dto.UserInfo).ID
	followId, err := strconv.Atoi(ctx.PostForm("followId"))
	if err != nil {
		result.SimpleErrorMessage("解析followId失败")
		return
	}
	if err := r.relationService.AddFollow(userId, uint(followId)); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("关注成功")
}

func (r *relationController) CancelFollow(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.Keys["user"].(*dto.UserInfo).ID
	followId, err := strconv.Atoi(ctx.PostForm("followId"))
	if err != nil {
		result.SimpleErrorMessage("解析followId失败")
		return
	}
	if err := r.relationService.CancelFollow(userId, uint(followId)); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("取关成功")
}

func (r *relationController) GetFollowList(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.Keys["user"].(*dto.UserInfo).ID
	followList, err := r.relationService.GetFollowList(userId)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(followList)
}

func (r *relationController) GetFollowerList(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.Keys["user"].(*dto.UserInfo).ID
	followerList, err := r.relationService.GetFollowerList(userId)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(followerList)
}
