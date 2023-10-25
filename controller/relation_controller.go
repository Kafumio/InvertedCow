package controller

import (
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
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
	userId := ctx.GetInt("userId")
	userToId := ctx.GetInt("userToId")
	if userId == 0 || userToId == 0 {
		result.SimpleErrorMessage("解析 userId 或 userToId 失败")
		return
	}
	if err := r.relationService.AddFollow(userId, userToId); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("关注成功")
}

func (r *relationController) CancelFollow(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.GetInt("userId")
	userToId := ctx.GetInt("userToId")
	if userId == 0 || userToId == 0 {
		result.SimpleErrorMessage("解析 userId 或 userToId 失败")
		return
	}
	if err := r.relationService.CancelFollow(userId, userToId); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("取关成功")
}

func (r *relationController) GetFollowList(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.GetInt("userId")
	if userId == 0 {
		result.SimpleErrorMessage("解析 userId 失败")
		return
	}
	followList, err := r.relationService.GetFollowList(userId)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(followList)
}

func (r *relationController) GetFollowerList(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	userId := ctx.GetInt("userId")
	if userId == 0 {
		result.SimpleErrorMessage("解析 userId 失败")
		return
	}
	followerList, err := r.relationService.GetFollowerList(userId)
	if err != nil {
		result.Error(err)
		return
	}
	result.SuccessData(followerList)
}
