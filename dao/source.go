package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

type SourceDao interface {
	InsertSource(db *gorm.DB, Source *po.Source) error
}

type sourceDao struct {
}

func NewSourceDao() SourceDao {
	return &sourceDao{}
}

func (p *sourceDao) InsertSource(db *gorm.DB, Source *po.Source) error {
	return db.Create(Source).Error
}
