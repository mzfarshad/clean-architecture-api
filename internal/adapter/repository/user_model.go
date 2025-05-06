package repository

import (
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"unique"`
	PasswordHash   string
	InactiveReason string
	Type           user.Type `gorm:"type:user_type;default:customer"`
	Status         bool      `gorm:"default:true"`
}

func mapUserToEntity(m *User) *user.Entity {
	return &user.Entity{
		Entity:         gormModelToDomainEntity(m.Model),
		Name:           m.Name,
		Email:          m.Email,
		PasswordHash:   m.PasswordHash,
		Type:           m.Type,
		InactiveReason: m.InactiveReason,
		Status:         m.Status,
	}
}
