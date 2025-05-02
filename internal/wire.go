//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/mzfarshad/music_store_api/internal/adapter/repository"
	"github.com/mzfarshad/music_store_api/internal/application"
)

func InjectDependencies() (*application.Container, error) {
	wire.Build(
		// basic dependencies

		// repositories
		repository.NewPostgresConnection,
		repository.NewUserRepo,
		// application's share services

		// application container
		application.NewContainer,
		application.NewAuthService,
		application.NewAdminService,
		application.NewCustomerService,
	)

	return &application.Container{}, nil
}
