package dto

import "InvertedCow/model/po"

type CommentTree struct {
	CommentId uint                   `json:"commentId"` // 评论ID
	Author    map[string]interface{} `json:"author"`    // 评论者
	Content   string                 `json:"content"`   // 树洞内容
	CreatedAt string                 `json:"createdAt"` // 评论时间
	LikeNum   int                    `json:"likeNum"`   // 点赞数
	IsLike    bool                   `json:"is_like"`   // 点赞或不点赞
	Children  []*CommentTree         `json:"children"`  // 子评论
}

type AddCommentForm struct {
	PostId   uint   `json:"postId"`
	UserId   uint   `json:"userId"`
	ParentId uint   `json:"parentId"`
	Content  string `json:"content"`
}

func NewCommentTree(comment *po.Comment) *CommentTree {
	return &CommentTree{
		CommentId: comment.ID,
		Author:    make(map[string]interface{}),
		Content:   comment.Content,
		LikeNum:   comment.LikeNum,
		IsLike:    false,
		Children:  make([]*CommentTree, 0),
	}
}
