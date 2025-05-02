package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/mzfarshad/music_store_api/config"
	"github.com/mzfarshad/music_store_api/internal/application"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/rest"
	v1 "github.com/mzfarshad/music_store_api/rest/api/v1"
	"github.com/mzfarshad/music_store_api/rest/api/v1/dashboard"
	"github.com/mzfarshad/music_store_api/rest/middleware"
)

func Serve(container *application.Container) error {
	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024, // 4 MB
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			// Retrieve the custom status code if it's a fiber.*Error
			var fiberErr *fiber.Error
			if ok := errors.As(err, &fiberErr); ok {
				code = fiberErr.Code
			}
			switch code {
			case 405:
				//	return presenter.HandleAppErrorResponse(ctx, &errorList, err.Error(), code, err.Error(), code)
				return rest.NewFailed(err).Handle(ctx)
			case 404:
				//	return presenter.HandleAppErrorResponse(ctx, &errorList, err.Error(), code, err.Error(), code)
				return rest.NewFailed(errs.New(errs.NotFound, err.Error())).Handle(ctx)
			default:
				// Return from handler
				return nil // (Copied form Nexu backend)
				//return ctx.Next() //??
			}
		},
	})

	app.Use(
		cors.New(),
		fiberLogger.New(fiberLogger.Config{Format: "[${ip}]:${port} ${status} - ${method} ${path}\n"}),
		skip.New(middleware.Secure(), middleware.ExcludeAccessSecuredRoutes),
		middleware.Recover(),
	)

	apiV1 := app.Group("/api/v1")
	// Register Customers APIs
	v1.Route(apiV1, container)
	dashboard.Route(apiV1, container)
	//logs.Info(context.Background(), fmt.Sprintf("Successfully initialized in %q environment.", config.Get().App.Env))

	return app.Listen(fmt.Sprintf(":%s", config.Get().App.Port))
}
