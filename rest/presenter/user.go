package presenter

import (
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
)

type User struct {
	rest.DTO
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	InactiveReason string `json:"inactive_reason"`
	Status         bool   `json:"status"`
}

func NewUser(entity *user.Entity) *User {
	return &User{
		Id:             entity.Id,
		Name:           entity.Name,
		Email:          entity.Email,
		InactiveReason: entity.InactiveReason,
		Status:         entity.Active,
	}
}
