package middleware

import (
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/rest"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Secure() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeaders := ctx.GetReqHeaders()["Authorization"]
		if authHeaders == nil || len(authHeaders) == 0 {
			err := errs.New(errs.Unauthorized, "Authorization header is required")
			return rest.NewFailed(err).Handle(ctx)
		}
		if !strings.HasPrefix(authHeaders[0], "Bearer ") {
			err := errs.New(errs.Unauthorized, "invalid bearer token")
			return rest.NewFailed(err).Handle(ctx)
		}
		claims, err := auth.ValidateToken(ctx.Context(), strings.TrimPrefix(authHeaders[0], "Bearer "))
		if err != nil {
			return rest.NewFailed(errs.New(errs.Unauthorized, "invalid claims").CausedBy(err)).Handle(ctx)
		}
		if claims == nil {
			return rest.NewFailed(errs.New(errs.Unauthorized, "claims is empty")).Handle(ctx)
		}
		ctx.Context().SetUserValue(auth.CtxKeyAuthUser, claims)
		return ctx.Next()
	}
}

func ExcludeAccessSecuredRoutes(ctx *fiber.Ctx) bool {
	uriPath := uri(ctx.Request().URI().Path())
	for _, route := range skippedAccessSecuredRoutes {
		if uriPath == route || route.IsSame(uriPath) {
			return true // skip the uri
		}
	}
	return false
}

var skippedAccessSecuredRoutes = []uri{
	"/api/v1/auth/signup",
}
