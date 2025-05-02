package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/rest"
	"github.com/mzfarshad/music_store_api/rest/presenter"
)

func authRouter(apiV1 fiber.Router, authService auth.CustomerUseCase) {
	v1Auth := apiV1.Group("/auth")
	v1Auth.Post("/signing", signingHandler(authService))
	v1Auth.Post("/signup", signupHandler(authService))
}

type signup struct {
	rest.DTO `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"  validate:"required"`    // TODO: read about email validation in golang validator
	Password string `json:"password"  validate:"required"` // TODO: read about password validation rules
}

func signupHandler(authService auth.CustomerUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[signup]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		tokens, err := authService.Signup(ctx.Context(), input.Name, input.Email, input.Password)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewAuthToken(tokens.Access)).Handle(ctx)
	}
}

type signing struct {
	rest.DTO `json:"-"`
	Email    string `json:"email"  validate:"required"`    // TODO: read about email validation in golang validator
	Password string `json:"password"  validate:"required"` // TODO: read about password validation rules
}

func signingHandler(authService auth.CustomerUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		input, err := rest.Request[signing]{}.Parse(ctx)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		tokens, err := authService.SignIn(ctx.Context(), input.Email, input.Password)
		if err != nil {
			return rest.NewFailed(err).Handle(ctx)
		}
		return rest.NewSuccess(presenter.NewAuthToken(tokens.Access)).Handle(ctx)
	}
}
