package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
	"github.com/mzfarshad/music_store_api/rest/presenter"
)

func usersRouter(v1Dashboard fiber.Router, userService user.AdminUseCase) {
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
		return rest.NewSuccess(nil).Handle(ctx)
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
		return rest.NewSuccess(nil).Handle(ctx)
	}
}

type updateProfile struct {
	rest.DTO `json:"_"`
	Name     string `json:"name" validate:"required"`
}

func updateMyProfile(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[updateProfile]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		if err = userService.UpdateMyProfile(ctx.Context(), input.Name); err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(nil).Handle(ctx)
	}
}

func searchInUsers(userService user.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		pagination, err := rest.NewPagination[user.SearchParams](ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		users, err := userService.SearchInUsers(ctx.Context(), pagination)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(
			rest.NewList(users, presenter.NewUser),
		).Paginate(pagination).Handle(ctx)
	}
}
