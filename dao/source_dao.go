package dao

import (
	"InvertedCow/model/po"
	"gorm.io/gorm"
)

// SourceDao
// TODO: crud
type SourceDao interface {
	InsertSource(db *gorm.DB, Source *po.Source) error
	GetSourceById(db *gorm.DB, sourceId uint) (*po.Source, error)
	GetSourceByPostId(db *gorm.DB, postId uint) (*po.Source, error)
}

type sourceDao struct {
}

func NewSourceDao() SourceDao {
	return &sourceDao{}
}

func (s *sourceDao) InsertSource(db *gorm.DB, Source *po.Source) error {
	return db.Create(Source).Error
}

func (s *sourceDao) GetSourceById(db *gorm.DB, sourceId uint) (*po.Source, error) {
	var source po.Source
	err := db.First(&source, sourceId).Error
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (s *sourceDao) GetSourceByPostId(db *gorm.DB, postId uint) (*po.Source, error) {
	var source po.Source
	err := db.Where("post_id = ?", postId).Find(&source).Error
	if err != nil {
		return nil, err
	}
	if source.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &source, nil
}
