//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/mzfarshad/music_store_api/internal/adapter/repository"
	"github.com/mzfarshad/music_store_api/internal/application"
	"github.com/mzfarshad/music_store_api/internal/application/admin"
	admin2 "github.com/mzfarshad/music_store_api/internal/application/cli/admin_cmd"
	"github.com/mzfarshad/music_store_api/internal/application/customer"
)

func NewContainer() (*application.Container, error) {
	wire.Build(
		// basic dependencies

		// repositories
		repository.NewPostgresConnection,
		repository.NewUserRepo,
		// share services

		// customer services
		customer.NewAuthService,
		customer.NewUserService,

		// admin services
		admin.NewUserService,
		admin.NewAuthService,

		// cli service
		admin2.NewAdminCmdService,

		// application container
		application.NewContainer,
	)

	return &application.Container{}, nil
}
