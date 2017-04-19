package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string
	Emails []Email
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}
