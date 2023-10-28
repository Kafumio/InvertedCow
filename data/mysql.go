package data

import (
	"InvertedCow/config"
	"InvertedCow/model/po"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// NewGormClient 初始化数据库配置
func NewGormClient(conf *config.AppConfig) *gorm.DB {
	cfg := conf.MySqlConfig
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 模型绑定
	err = db.AutoMigrate(
		po.User{},
		po.Post{},
		po.Source{},
	)
	if err != nil {
		log.Println(err)
	}

	return db
}
