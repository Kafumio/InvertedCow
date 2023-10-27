package po

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	PostID   uint   `gorm:"column:uid;comment:原动态 业务唯一标识" json:"post_id"`
	FileName string `gorm:"column:file_name;comment:存储方标识" json:"file_name"`
	Hash     string `gorm:"column:hash" json:"hash"`
	Size     int64  `gorm:"column:size;comment:文件大小" json:"size"`
	Key      string `gorm:"column:key;comment:key" json:"key"`
	Bucket   string `gorm:"column:bucket;" json:"bucket"`
}
