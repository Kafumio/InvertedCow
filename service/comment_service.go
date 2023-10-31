package service

import (
	"InvertedCow/dao"
	"InvertedCow/data"
	e "InvertedCow/error"
	"InvertedCow/model/dto"
	"InvertedCow/model/po"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type CommentService interface {
	// AddComment 添加评论
	AddComment(comment *po.Comment) *e.Error
	// DeleteComment 删除评论
	DeleteComment(commentId, postId uint) *e.Error
	// LikeComment 点赞评论
	LikeComment(commentId uint) *e.Error
	// DisLikeComment 取消点赞评论
	DisLikeComment(commentId uint) *e.Error
	// GetPostComments 获取动态的评论树数组
	GetPostComments(postId uint) ([]dto.CommentTree, *e.Error)
	// GetPostCommentChild 获取动态评论树的子评论
	GetPostCommentChild(parentId uint, commentTree *dto.CommentTree) *e.Error
}

type commentService struct {
	db    *gorm.DB
	cos   *data.Cos
	redis *redis.Client
	cd    dao.CommentDao
	ud    dao.UserDao
	pd    dao.PostDao
}

func NewCommentService(db *gorm.DB, cos *data.Cos, redis *redis.Client, cd dao.CommentDao, ud dao.UserDao, pd dao.PostDao) CommentService {
	return &commentService{
		db:    db,
		cos:   cos,
		redis: redis,
		cd:    cd,
		ud:    ud,
		pd:    pd,
	}
}

func (c *commentService) AddComment(comment *po.Comment) *e.Error {
	if err := c.cd.AddComment(c.db, comment); err != nil {
		return e.ErrAddCommentFailed
	}
	post, _ := c.pd.GetPostByID(c.db, comment.PostId)
	post.CommentNum = post.CommentNum + 1
	if err := c.pd.UpdatePost(c.db, post); err != nil {
		return e.ErrMysql
	}
	return nil
}

func (c *commentService) DeleteComment(commentId, postId uint) *e.Error {
	if err := c.cd.DeleteComment(c.db, commentId); err != nil {
		return e.ErrDeleteCommentFailed
	}
	post, _ := c.pd.GetPostByID(c.db, postId)
	post.CommentNum = post.CommentNum - 1
	if err := c.pd.UpdatePost(c.db, post); err != nil {
		return e.ErrMysql
	}
	return nil
}

func (c *commentService) LikeComment(commentId uint) *e.Error {
	if err := c.cd.LikeComment(c.db, commentId); err != nil {
		return e.ErrMysql
	}
	return nil
}

func (c *commentService) DisLikeComment(commentId uint) *e.Error {
	if err := c.cd.DisLikeComment(c.db, commentId); err != nil {
		return e.ErrMysql
	}
	return nil
}

func (c *commentService) GetPostComments(postId uint) ([]dto.CommentTree, *e.Error) {
	var commentTrees []dto.CommentTree
	comments, err := c.cd.GetCommentsByPostId(c.db, postId)
	if err != nil {
		return commentTrees, e.ErrMysql
	}
	for _, comment := range comments {
		var user *po.User
		cid := comment.ID
		uid := comment.UserId
		user, _ = c.ud.GetUserByID(c.db, uid)
		commentTree := dto.CommentTree{
			CommentId: cid,
			Author:    gin.H{"uid": uid, "username": user.Username, "avatar": user.Avatar},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04"),
			LikeNum:   comment.LikeNum,
			IsLike:    false,
			Children:  []*dto.CommentTree{},
		}
		err := c.GetPostCommentChild(cid, &commentTree)
		if err != nil {
			return commentTrees, err
		}
		commentTrees = append(commentTrees, commentTree)
	}
	return commentTrees, nil
}

func (c *commentService) GetPostCommentChild(parentId uint, commentTree *dto.CommentTree) *e.Error {
	comments, err := c.cd.GetCommentsByParentId(c.db, parentId)
	if err != nil {
		return e.ErrMysql
	}
	// 查询二级及以下的多级评论
	for i, _ := range comments {
		var user *po.User
		cid := comments[i].ID
		uid := comments[i].UserId
		user, err = c.ud.GetUserByID(c.db, uid)
		if err != nil {
			return e.ErrMysql
		}
		child := dto.CommentTree{
			CommentId: cid,
			Author:    gin.H{"uid": user.ID, "username": user.Username, "avatar": user.Avatar},
			Content:   comments[i].Content,
			CreatedAt: comments[i].CreatedAt.Format("2006-01-02 15:04"),
			LikeNum:   comments[i].LikeNum,
			IsLike:    false,
			Children:  []*dto.CommentTree{},
		}
		commentTree.Children = append(commentTree.Children, &child)
		c.GetPostCommentChild(cid, &child)
	}
	return nil
}
