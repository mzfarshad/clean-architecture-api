package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/config"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/rest"
)

const recoverMsg = "An error has occurred. Please try again in a few minutes"

func Recover() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if config.Get().App.Recover {
			defer func() {
				if r := recover(); r != nil {
					// TODO: Set Layer
					//logCtx := logs.SetLayer(ctx.Context(), logs.LayerMiddleware)
					// TODO: Set URL
					//pathUrl := ctx.Path()
					//logCtx = logs.SetUrl(logCtx, pathUrl)

					//ip := ctx.Get("Cf-Connecting-Ip", "")
					//logCtx = logs.SetClientIP(logCtx, ip)
					// TODO: Set HTTP request
					//logCtx = logs.SetHttpRequest(logCtx, ctx.Body())
					// TODO: Set UserId
					//claims, ok := ctx.Context().UserValue(auth.CtxKeyAuthUser).(*auth.UserClaims)
					//if ok && claims != nil {
					//	logCtx = logs.SetUserId(logCtx, claims.UserId)
					//}
					//logs.Error(logCtx, fmt.Sprintf("Panic: %s  |  stack trace : %s", r, string(debug.Stack())))
					_ = rest.NewFailed(errs.New(errs.Internal, recoverMsg)).Handle(ctx)
				}
			}()
		}
		return ctx.Next()
	}
}
