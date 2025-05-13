package presenter

import (
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
)

func NewSuccessResponse(id uint, msg string) *SuccessResponse {
	return &SuccessResponse{
		Message: msg,
		UserId:  id,
	}
}

type SuccessResponse struct {
	rest.DTO `json:"_"`
	Message  string `json:"message,omitempty"`
	UserId   uint   `json:"user_id,omitempty"`
}

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
		Status:         entity.Status,
	}
}
