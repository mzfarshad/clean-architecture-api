package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/application"
)

func Route(apiV1 fiber.Router, container *application.Container) {
	authRouter(apiV1, container.Customer.AuthService)
}
