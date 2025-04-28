package repository

import (
	"github.com/mzfarshad/music_store_api/infra/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"unique"`
	PasswordHash   string
	InactiveReason string
	Type           user.Type `gorm:"type:user_type;default:'customer'"` // TODO: make an enum in database using a common function
	Status         bool      `gorm:"default:true"`
}

func mapUserToEntity(m *User) *user.Entity {
	return &user.Entity{
		Entity:         gormModelToDomainEntity(m.Model),
		Name:           m.Name,
		Email:          m.Email,
		Type:           m.Type,
		InactiveReason: m.InactiveReason,
		Status:         m.Status,
	}
}

// must be defined int handler layer
// type UserSearch struct {
// 	Name  string `form:"name"`
// 	Email string `form:"email"`
// 	Page  int    `form:"page"`
// 	Limit int    `form:"limit"`
// }
