package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Setup() error {
	DB, err := gorm.Open(sqlite.Open("./sqlite/exater.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(User{})
	if err != nil {
		return err
	}

	return nil
}
