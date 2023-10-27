package po

import (
	"gorm.io/gorm"
)

// Source 代表一个资源，这里可以特定为短视频
// 目前的设计中，Source与Post是一对一的关系
// 在发布动态的过程中，Post会先于Source创建，它们的关联会在七牛云的回调请求中进行设定
// 简化为：1. 更改Post状态为发布成功 2. Source创建时关联特定的Post
type Source struct {
	gorm.Model
	PostID   uint   `gorm:"column:uid;comment:原动态 业务唯一标识" json:"post_id"`
	FileName string `gorm:"column:file_name;comment:存储方标识" json:"file_name"`
	Hash     string `gorm:"column:hash" json:"hash"`
	Size     int64  `gorm:"column:size;comment:文件大小" json:"size"`
	Key      string `gorm:"column:key;comment:key" json:"key"`
	Bucket   string `gorm:"column:bucket;" json:"bucket"`
}
