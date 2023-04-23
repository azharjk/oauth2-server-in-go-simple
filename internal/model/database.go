package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	t, err := gorm.Open(sqlite.Open("./sqlite/exater.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	DB = t

	return nil
}

func Migrate() error {
	err := Connect()
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(User{}, Client{})
	if err != nil {
		return err
	}

	return nil
}
