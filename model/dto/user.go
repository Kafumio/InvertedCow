package dto

import "InvertedCow/model/po"

type UserInfo struct {
	ID        uint   `json:"id"`
	Avatar    string `json:"avatar"`
	LoginName string `json:"loginName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func NewUserInfo(user *po.User) *UserInfo {
	userInfo := &UserInfo{
		ID:        user.ID,
		Avatar:    user.Avatar,
		LoginName: user.LoginName,
		Username:  user.Username,
		Email:     user.Email,
	}
	return userInfo
}
