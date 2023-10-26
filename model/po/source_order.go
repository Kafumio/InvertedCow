package po

import "gorm.io/gorm"

type SourceOrder struct {
	gorm.Model
	Seq      int   `gorm:"column:seq;comment:次序" json:"seq"`
	OriginId int64 `gorm:"column:origin_id;comment:动态id" json:"origin_id"`
	SourceId int64 `gorm:"column:source_id;comment:资源id" json:"source_id"`
}
