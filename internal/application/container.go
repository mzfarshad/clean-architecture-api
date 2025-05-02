package application

import (
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

type share struct {
	// For future when we need to use a service in another service preventing import cycle
}

type customer struct {
	AuthService auth.CustomerUseCase
	UserService user.CustomerUseCase
}

type admin struct {
	UserService user.AdminUseCase
}

type Container struct {
	Share    share
	Admin    admin
	Customer customer
}

func NewContainer(
	// Share Services

	// Admin Services
	adminUserService user.AdminUseCase,

	// Customer Services
	authService auth.CustomerUseCase,
	userService user.CustomerUseCase,

	// Other Use Cases

) *Container {
	return &Container{
		Share: share{},
		Admin: admin{
			UserService: adminUserService,
		},
		Customer: customer{
			AuthService: authService,
			UserService: userService,
		},
	}
}
