package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Setup() {
	DB, _ = gorm.Open(sqlite.Open("./sqlite/exater.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	_ = DB.AutoMigrate(User{})
}
