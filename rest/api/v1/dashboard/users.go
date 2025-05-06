package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
)

func usersRouter(v1Dashboard fiber.Router, userService user.AdminUseCase) {
	// TODO
	users := v1Dashboard.Group("/users")
	users.Get("", searchInUsers(userService))
	users.Put("/updateMyProfile", updateMyProfile(userService))
	users.Put("/deactivate/:id", deactivateUser(userService))
	users.Put("/reactivate/:id", reactivateUser(userService))
}

type deactiveUser struct {
	rest.DTO `json:"-"`
	Id       uint   `params:"id" validate:"required"` // TODO: Test this
	Reason   string `json:"reason" validate:"required"`
}

func deactivateUser(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[deactiveUser]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.DeactivateUser(ctx.Context(), input.Id, input.Reason); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return nil
	}
}

type reactiveUser struct {
	rest.DTO `json:"_"`
	Id       uint `params:"id" validate:"required"`
}

func reactivateUser(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[reactiveUser]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.ReactivateUser(ctx.Context(), input.Id); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return nil
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
		return nil
	}
}

type searchUsers struct {
	rest.DTO `json:"_"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
}

type pagination struct {
	rest.DTO       `json:"_"`
	userPagination user.PaginationParams
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

		users, err := userService.SearchInUsers(ctx.Context(), searchParam)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		var usersPage pagination
		usersPage.userPagination.TotalPages = users.TotalPages
		usersPage.userPagination.TotalData = users.TotalData
		usersPage.userPagination.Result = users.Result
		return rest.NewSuccess(usersPage).Handle(ctx)
	}
}
