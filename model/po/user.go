package po

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Avatar       string `gorm:"column:avatar" json:"avatar"`
	Username     string `gorm:"column:username" json:"username"`
	LoginName    string `gorm:"column:login_name" json:"loginName"`
	Password     string `gorm:"column:password" json:""`
	Salt         string `gorm:"column:salt" json:""`
	Email        string `gorm:"column:email" json:"email"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Introduction string `gorm:"column:introduction" json:"introduction"`
	// 1表示男 2表示女
	Sex      int       `gorm:"column:sex" json:"sex"`
	BirthDay time.Time `gorm:"column:birth_day" json:"birthDay"`
}
