package dto

import "InvertedCow/model/po"

type PostDtoForView struct {
	ID uint `json:"id"`
	// Text 正文
	Text string `json:"text"`
	// VideoUrl 视频url
	VideoUrl string `json:"videoUrl"`
	// Publisher 发布者id
	Publisher uint `json:"publisher"`
	// PublisherAvatar 发布者头像
	PublisherAvatar string `json:"publisherAvatar"`
	// CommentNum 评论数
	CommentNum int `json:"commentNum"`
}

func NewPostDtoForView(post *po.Post) *PostDtoForView {
	return &PostDtoForView{
		ID:         post.ID,
		Text:       post.Text,
		Publisher:  post.Publisher,
		CommentNum: post.CommentNum,
	}
}
