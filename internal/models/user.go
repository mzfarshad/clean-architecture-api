package models

import "gorm.io/gorm"

type UserType string

const (
	UserTypeUser  UserType = "user"
	UserTypeAdmin UserType = "admin"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Type     UserType `gorm:"default:user"`
}
