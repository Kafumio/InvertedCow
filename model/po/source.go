package po

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	Hash string `gorm:"column:hash" json:"hash"`
	Size int64  `gorm:"column:size;comment:文件大小" json:"size"`
	Key  string `gorm:"column:source_url;comment:资源路径" json:"key"`
}
