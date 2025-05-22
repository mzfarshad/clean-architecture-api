package repository

import (
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) user.Repository {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}

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
	entity := &user.Entity{
		Name:           m.Name,
		Email:          m.Email,
		InactiveReason: m.InactiveReason,
		Type:           m.Type,
		Status:         m.Status,
	}
	entity.SetPasswordHash(m.PasswordHash)
	return entity
}
