package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/application"
	"github.com/mzfarshad/music_store_api/internal/domain"
	"github.com/mzfarshad/music_store_api/rest/middleware"
)

func Route(apiV1 fiber.Router, container *application.Container) {
	v1Dashboard := apiV1.Group("/dashboard", middleware.Only(domain.Admin))

	usersRouter(v1Dashboard, container.Admin.UserService)
	authRouter(apiV1, container.Admin.AuthService)
}
