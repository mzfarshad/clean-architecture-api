package presenter

import (
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
)

func NewDashboardResponse(id uint, msg string) *DashboardResponse {
	return &DashboardResponse{
		Message: msg,
		UserId:  id,
	}
}

type DashboardResponse struct {
	rest.DTO `json:"_"`
	Message  string `json:"message,omitempty"`
	UserId   uint   `json:"user_id,omitempty"`
}

type SearchUserDto struct {
	rest.DTO
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	InactiveReason string `json:"inactive_reason"`
	Status         bool   `json:"status"`
}

func MapUserEntityToSearchUserDTO(entity *user.Entity) *SearchUserDto {
	return &SearchUserDto{
		Id:             entity.Id,
		Name:           entity.Name,
		Email:          entity.Email,
		InactiveReason: entity.InactiveReason,
		Status:         entity.Status,
	}
}
