package model

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Avatar   string `json:"avatar"`
	Username string `json:"username"`

}

type Tweet struct{
	gorm.Model
	Tweet string `json:"tweet"`
}