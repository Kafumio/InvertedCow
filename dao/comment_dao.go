package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type CommentDao interface {
	// AddComment 添加评论
	AddComment(db *gorm.DB, comment *po.Comment) error
	// DeleteComment 删除评论(由于设置级联删除，子评论也会被一并删除)
	DeleteComment(db *gorm.DB, commentId uint) error
	// LikeComment 点赞评论
	LikeComment(db *gorm.DB, commentId uint) error
	// DisLikeComment 取消点赞该评论
	DisLikeComment(db *gorm.DB, commentId uint) error
	// GetCommentsByPostId 通过动态id获取该动态的所有评论
	GetCommentsByPostId(db *gorm.DB, postId uint) ([]po.Comment, error)
	// GetCommentsByParentId 通过父id获取该评论的所有子评论
	GetCommentsByParentId(db *gorm.DB, parentId uint) ([]po.Comment, error)
}

type commentDao struct {
}

func NewCommentDao() CommentDao {
	return &commentDao{}
}

func (c *commentDao) AddComment(db *gorm.DB, comment *po.Comment) error {
	return db.Create(comment).Error
}

func (c *commentDao) DeleteComment(db *gorm.DB, commentId uint) error {
	return db.Delete(&po.Comment{}, commentId).Error
}

func (c *commentDao) LikeComment(db *gorm.DB, commentId uint) error {
	return db.Where("id = ?", commentId).UpdateColumn("like_num",
		gorm.Expr("like_num + ?", 1)).Error
}

func (c *commentDao) DisLikeComment(db *gorm.DB, commentId uint) error {
	return db.Where("id = ?", commentId).UpdateColumn("like_num",
		gorm.Expr("like_num - ?", 1)).Error
}

func (c *commentDao) GetCommentsByPostId(db *gorm.DB, postId uint) ([]po.Comment, error) {
	var comments []po.Comment
	err := db.Where("post_id = ? AND parent_id IS NULL", postId).Order("id desc").Find(&comments).Error
	return comments, err
}

func (c *commentDao) GetCommentsByParentId(db *gorm.DB, parentId uint) ([]po.Comment, error) {
	var comments []po.Comment
	err := db.Where("parent_id = ?", parentId).Find(&comments).Error
	return comments, err
}
