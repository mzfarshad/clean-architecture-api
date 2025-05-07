package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
	"github.com/mzfarshad/music_store_api/rest/presenter"
)

func usersRouter(v1Dashboard fiber.Router, userService user.AdminUseCase) {
	// TODO
	users := v1Dashboard.Group("/users")
	users.Get("", searchInUsers(userService))
	users.Put("/updateMyProfile", updateMyProfile(userService))
	users.Put("/deactivate/:id", deactivateUser(userService))
	users.Put("/reactivate/:id", reactivateUser(userService))
}

type deactivatedUser struct {
	rest.DTO `json:"-"`
	Id       uint   `params:"id" validate:"required"` // TODO: Test this
	Reason   string `json:"reason" validate:"required"`
}

func deactivateUser(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[deactivatedUser]{}.ParseParams(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.DeactivateUser(ctx.Context(), input.Id, input.Reason); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewDashboardResponse(input.Id, "User deactivated successfully")).Handle(ctx)
	}
}

type reactivatedUser struct {
	rest.DTO `json:"_"`
	Id       uint `params:"id" validate:"required"`
}

func reactivateUser(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[reactivatedUser]{}.ParseParams(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.ReactivateUser(ctx.Context(), input.Id); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewDashboardResponse(input.Id, "User reactivated successfully")).Handle(ctx)
	}
}

type updateProfile struct {
	rest.DTO `json:"_"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func updateMyProfile(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[updateProfile]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.UpdateMyProfile(ctx.Context(), input.Name, input.Email); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewDashboardResponse(0, "User updated successfully")).Handle(ctx)
	}
}

type searchUsers struct {
	rest.DTO `json:"_"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Limit    int    `json:"limit" form:"limit"`
	Page     int    `json:"page" form:"page"`
}

func searchInUsers(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[searchUsers]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		var searchParam user.SearchParams
		searchParam.Limit = input.Limit
		searchParam.Name = input.Name
		searchParam.Email = input.Email
		searchParam.Page = input.Page

		pagesData, err := userService.SearchInUsers(ctx.Context(), searchParam)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		dtoPagesData := rest.NewList(pagesData.Result, presenter.MapUserEntityToUserDTO)
		pagination := presenter.NewUserPaginationAdapter(searchParam, pagesData)
		return rest.NewSuccess(dtoPagesData).Paginate(pagination).Handle(ctx)
	}
}
