package dto

import (
	"InvertedCow/model/po"
	"time"
)

// UserInfo token里面存储的数据
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

// AccountInfo 用户读取账号信息时返回的对象
type AccountInfo struct {
	ID           uint   `json:"id"`
	Avatar       string `json:"avatar"`
	LoginName    string `json:"loginName"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Introduction string `json:"introduction"`
	// 1表示男 2表示女
	Sex      int       `json:"sex"`
	BirthDay time.Time `json:"birthDay"`
}

func NewAccountInfo(user *po.User) *AccountInfo {
	userInfo := &AccountInfo{
		ID:           user.ID,
		Avatar:       user.Avatar,
		LoginName:    user.LoginName,
		Username:     user.Username,
		Email:        user.Email,
		Introduction: user.Introduction,
		Sex:          user.Sex,
		BirthDay:     user.BirthDay,
	}
	return userInfo
}
