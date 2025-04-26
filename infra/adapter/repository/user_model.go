package repository

import (
	"github.com/mzfarshad/music_store_api/infra/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
	Type         user.Type `gorm:"default:user"` // TODO: make an enum in database using a common function
}

func mapUserToEntity(m *User) *user.Entity {
	return &user.Entity{
		Entity: gormModelToDomainEntity(m.Model),
		Name:   m.Name,
		Email:  m.Email,
		Type:   m.Type,
	}
}
