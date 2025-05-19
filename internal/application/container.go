package application

import (
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

type admin struct {
	UserService user.AdminUseCase
	AuthService auth.AdminUseCase
}

type customer struct {
	AuthService auth.CustomerUseCase
	UserService user.CustomerUseCase
}

type share struct {
	// For future when we need to use a service in another service preventing import cycle
}

type cli struct {
	// For future when we need to use a service in another service preventing import cycle
}

type Container struct {
	Admin    admin
	Customer customer
	Share    share
	Cli      cli
}

func NewContainer(
	// Share Services

	// Admin Services
	adminUserService user.AdminUseCase,
	adminAuthService auth.AdminUseCase,

	// Customer Services
	authService auth.CustomerUseCase,
	userService user.CustomerUseCase,

	// Cli service

) *Container {
	return &Container{
		Admin: admin{
			UserService: adminUserService,
			AuthService: adminAuthService,
		},
		Customer: customer{
			AuthService: authService,
			UserService: userService,
		},
		Share: share{},
		Cli:   cli{},
	}
}
