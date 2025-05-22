package user

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain"
	"github.com/mzfarshad/music_store_api/pkg/search"
)

type Repository interface {
	Query
	Command
}

type Query interface {
	First(context.Context, Where) (*Entity, error)
	Search(context.Context, *search.Pagination[SearchParams]) ([]*Entity, error)
}

type Command interface {
	Create(ctx context.Context, params CreateParams) (*Entity, error)
	Update(ctx context.Context, entity *Entity) error
}

type Where struct {
	Id    uint
	Type  Type
	Email string
}

type SearchParams struct {
	Name  string `query:"name"`
	Email string `query:"email"`
	Type  Type   `query:"type"`
}

type CreateParams struct {
	domain.Validatable
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Type     Type   `validate:"required,oneof:customer"`
}
