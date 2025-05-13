package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
	"github.com/mzfarshad/music_store_api/rest/presenter"
)

const (
	DefaultPageSize = 20
	DefaultPage     = 1
)

func usersRouter(v1Dashboard fiber.Router, userService user.AdminUseCase) {
	// TODO
	users := v1Dashboard.Group("/users")
	users.Get("", searchInUsers(userService))
	users.Put("/updateMyProfile", updateMyProfile(userService))
	users.Put("/deactivate/:id", deactivateUser(userService))
	users.Put("/reactivate/:id", reactivateUser(userService))
}

type deactivatedUserId struct {
	rest.DTO `json:"-"`
	Id       uint `params:"id" validate:"required"` // TODO: Test this
}
type deactivatedUserReason struct {
	rest.DTO `json:"_"`
	Reason   string `json:"reason" validate:"required"`
}

func deactivateUser(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := rest.Request[deactivatedUserId]{}.ParseParams(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		reason, err := rest.Request[deactivatedUserReason]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err := userService.DeactivateUser(ctx.Context(), userId.Id, reason.Reason); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewSuccessResponse(userId.Id, "User deactivated successfully")).Handle(ctx)
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
		return rest.NewSuccess(presenter.NewSuccessResponse(input.Id, "User reactivated successfully")).Handle(ctx)
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
		return rest.NewSuccess(presenter.NewSuccessResponse(0, "User updated successfully")).Handle(ctx)
	}
}

func searchInUsers(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		pagination, err := rest.NewPagination[user.SearchParams](ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}

		if pagination.Size() < 1 {
			pagination.Query.Limit = DefaultPageSize
		}
		if pagination.Page() < 1 {
			pagination.Query.Page = DefaultPage
		}

		pagesData, err := userService.SearchInUsers(ctx.Context(), pagination.Query)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}

		dtoPagesData := rest.NewList(pagesData.Result, presenter.NewUser)
		pagination.WithTotal(int64(pagesData.TotalData))

		return rest.NewSuccess(dtoPagesData).Paginate(pagination).Handle(ctx)
	}
}
