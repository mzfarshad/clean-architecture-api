package user

import (
	"context"
	"github.com/mzfarshad/music_store_api/pkg/search"
)

// AdminUseCase represents what an TypeAdmin can do with user Entity.
type AdminUseCase interface {
	ReactivateUser(ctx context.Context, userId uint) (*Entity, error)
	DeactivateUser(ctx context.Context, userId uint, reason string) (*Entity, error)
	SearchInUsers(context.Context, *search.Pagination[SearchParams]) ([]*Entity, error)
	UpdateMyProfile(ctx context.Context, name string) (*Entity, error)
}

// CustomerUseCase represents what a TypeCustomer can do with user Entity.
type CustomerUseCase interface {
	UpdateMyName(ctx context.Context, name string) (*Entity, error)
}
