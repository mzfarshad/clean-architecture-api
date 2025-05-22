package user

import (
	"context"
)

type Repository interface {
	Query
	Command
}

type Query interface {
	First(context.Context, Where) (*Entity, error)
	Find(context.Context, SearchParams) (*PaginationParams, error)
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
