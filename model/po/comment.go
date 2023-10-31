package po

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostId   uint     `gorm:"column:not null;index" json:"postId"`
	Post     Post     `gorm:"foreignKey:PostId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId   uint     `gorm:"not null;index" json:"userId"`
	User     User     `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ParentId *uint    `gorm:"index;default:NULL" json:"parentId"`
	Parent   *Comment `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content  string   `gorm:"type:varchar(255);not null" json:"content"`
	LikeNum  int      `gorm:"column:like_num;not null" json:"likeNum"`
}
