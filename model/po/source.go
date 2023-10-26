package po

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	UID        string `gorm:"column:uid;comment:业务唯一标识" json:"uid"`
	Hash       string `gorm:"column:hash" json:"hash"`
	Size       int64  `gorm:"column:size;comment:文件大小" json:"size"`
	SourceUrl  string `gorm:"column:source_url;comment:资源链接" json:"source_url"`
	SourceType int    `gorm:"column:source_type;comment:资源类型 1图片 2视频" json:"source_type"`
	OriginId   int64  `gorm:"column:origin_id;comment:原动态id" json:"origin_id"`
	IsDeleted  int    `gorm:"column:is_deleted;comment:删除标志位 1表示未删除 2表示已删除" json:"is_deleted"`
}
