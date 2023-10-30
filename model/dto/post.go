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
	// IsFollowed 是否已关注发布者
	IsFollowed bool `json:"isFollowed"`
	// LikedCount 点赞注数量
	LikedCount int `json:"likeCount"`
	// IsLiked 用户是否点赞
	IsLiked bool `json:"isLiked"`
}

func NewPostDtoForView(post *po.Post) *PostDtoForView {
	return &PostDtoForView{
		ID:         post.ID,
		Text:       post.Text,
		Publisher:  post.Publisher,
		LikedCount: post.LikedCount,
	}
}
