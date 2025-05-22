package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/rest"
)

func Only(userTypes ...domain.UserType) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userClaims, ok := ctx.Context().UserValue(auth.CtxKeyAuthUser).(*auth.UserClaims)
		if !ok {
			return rest.NewFailed(errs.New(errs.Unauthorized, "Unauthenticated user")).Handle(ctx)
		}
		if len(userTypes) > 0 {
			if !userClaims.UserType.Is(userTypes[0], userTypes[1:]...) {
				return rest.NewFailed(errs.New(errs.Forbidden, fmt.Sprintf("Access denied for a %s", userClaims.UserType))).Handle(ctx)
			}
		}

		return ctx.Next()
	}
}
