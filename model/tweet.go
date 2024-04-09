package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username"`
	Avatar   string  `json:"avatar"`
	Tweets   []Tweet `gorm:"foreignKey:UserID"`
}

type Tweet struct {
	gorm.Model
	Tweet  string `json:"tweet"`
	UserID uint   `json:"userId"`
	User   User   `json:"user"`
}
