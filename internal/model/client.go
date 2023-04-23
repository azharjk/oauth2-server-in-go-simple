package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Secret      string `gorm:"notNull"`
	Type        string `gorm:"notNull"`
	RedirectURI string `gorm:"notNull"`
}
