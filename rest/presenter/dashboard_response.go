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

type UserDto struct {
	rest.DTO
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	InactiveReason string `json:"inactive_reason"`
	Status         bool   `json:"status"`
}

func MapUserEntityToUserDTO(entity *user.Entity) *UserDto {
	return &UserDto{
		Id:             entity.Id,
		Name:           entity.Name,
		Email:          entity.Email,
		InactiveReason: entity.InactiveReason,
		Status:         entity.Status,
	}
}

func NewUserPaginationAdapter(params user.SearchParams, paginationParams *user.PaginationParams) *UserPaginationAdapter {
	filters := make(map[string]any)
	if params.Name != "" {
		filters["name"] = params.Name
	}
	if params.Email != "" {
		filters["email"] = params.Email
	}
	return &UserPaginationAdapter{
		Data:     paginationParams,
		Page_:    params.Page,
		PageSize: params.Limit,
		Filters_: filters,
	}
}

type UserPaginationAdapter struct {
	Data     *user.PaginationParams
	Page_    int
	PageSize int
	Filters_ map[string]any
}

func (p UserPaginationAdapter) Size() int {
	return p.PageSize
}

func (p UserPaginationAdapter) Page() int {
	return p.Page_
}

func (p UserPaginationAdapter) Total() int64 {
	return int64(p.Data.TotalData)
}

func (p UserPaginationAdapter) Filters() any {
	if len(p.Filters_) == 0 {
		return nil
	}
	return p.Filters
}
