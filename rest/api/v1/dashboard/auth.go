package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/rest"
	"github.com/mzfarshad/music_store_api/rest/presenter"
)

func authRouter(apiV1 fiber.Router, userService auth.AdminUseCase) {
	admin := apiV1.Group("/admin/login")
	admin.Post("", loginHandler(userService))
}

type login struct {
	rest.DTO `json:"-"`
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

func loginHandler(userService auth.AdminUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[login]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		tokens, err := userService.SingIn(ctx.Context(), input.Email, input.Password)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewAuthToken(tokens.Access)).Handle(ctx)
	}
}
