package controller

import (
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"InvertedCow/model/vo"
	"InvertedCow/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommentController interface {
	// AddComment 添加评论
	AddComment(ctx *gin.Context)
	// DeleteComment 删除评论
	DeleteComment(ctx *gin.Context)
	// LikeComment 点赞评论
	LikeComment(ctx *gin.Context)
	// DisLikeComment 取消点赞评论
	DisLikeComment(ctx *gin.Context)
	// GetPostComments 获取动态的所有评论
	GetPostComments(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentController{
		commentService: commentService,
	}
}

func (c *commentController) AddComment(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	addCommentForm := &dto.AddCommentForm{}
	if err := ctx.ShouldBindJSON(addCommentForm); err != nil {
		result.SimpleErrorMessage("解析添加评论表单失败")
		return
	}
	comment := &po.Comment{
		UserId:  addCommentForm.UserId,
		PostId:  addCommentForm.PostId,
		Content: addCommentForm.Content,
	}
	// 如果是第一条评论，也就是父评论，那么前端在传的时候ParentId为0就好，因为它不可能有父评论。
	// 因此通过判断ParentId是0后，就不需要传入该字段
	if addCommentForm.ParentId != 0 {
		comment.ParentId = &addCommentForm.ParentId
	}
	if err := c.commentService.AddComment(comment); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("发表评论成功")
}

func (c *commentController) DeleteComment(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	commentId, err := strconv.Atoi(ctx.PostForm("commentId"))
	if err != nil {
		result.SimpleErrorMessage("解析commentId失败")
		return
	}
	post, ok := ctx.Get("post")
	if !ok {
		result.Error(e.ErrBadRequest)
		return
	}
	postId := post.(*dto.PostDtoForView).ID
	if err := c.commentService.DeleteComment(uint(commentId), postId); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("删除评论成功")
}

func (c *commentController) LikeComment(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	commentId, err := strconv.Atoi(ctx.PostForm("commentId"))
	if err != nil {
		result.SimpleErrorMessage("解析commentId失败")
		return
	}
	if err := c.commentService.LikeComment(uint(commentId)); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("点赞评论成功")
}

func (c *commentController) DisLikeComment(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	commentId, err := strconv.Atoi(ctx.PostForm("commentId"))
	if err != nil {
		result.SimpleErrorMessage("解析commentId失败")
		return
	}
	if err := c.commentService.DisLikeComment(uint(commentId)); err != nil {
		result.Error(err)
		return
	}
	result.SuccessMessage("取消点赞评论成功")
}

func (c *commentController) GetPostComments(ctx *gin.Context) {
	result := vo.NewResult(ctx)
	postIdStr := ctx.Param("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		result.Error(e.ErrBadRequest)
		return
	}
	commentTrees, err2 := c.commentService.GetPostComments(uint(postId))
	if err != nil {
		result.Error(err2)
		return
	}
	result.SuccessData(commentTrees)
}
