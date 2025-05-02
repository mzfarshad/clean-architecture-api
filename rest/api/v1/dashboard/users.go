package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/rest"
)

func usersRouter(v1Dashboard fiber.Router, userService user.AdminUseCase) {
	// TODO
	//users := v1Dashboard.Group("/users")
	//users.Get("", searchInUsers())
	//users.Get("/:id", getUser())
	//users.Post("", createNewUser())
	//users.Put("/:id", updateUser())
	//users.Put("/:id/deactivate", deactivateUser())
	//users.Put("/:id/reactivate", reactivateUser())
}

type getUser struct {
	rest.DTO `json:"-"`
	Id       uint `params:"id" validate:"required"` // TODO: Test this
}

// func getUser() {
// rest.Request[getUser]{}.
//}
