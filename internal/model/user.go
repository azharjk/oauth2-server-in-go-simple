package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"notNull"`
	Email    string `gorm:"notNull;unique;uniqueIndex"`
	Password string `gorm:"notNull"`
}
