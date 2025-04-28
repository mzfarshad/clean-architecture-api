package user

import "context"

type Repository interface {
	Query
	Command
}

// TODO: @Farshad Search about CQRS

type Query interface { // GORM: First, Last, Take, Find
	FirstByEmail(ctx context.Context, email string) (*Entity, error)
	FirstById(ctx context.Context, id uint) (*Entity, error)
	Find(ctx context.Context, params SearchParams) (*ResponsePagination, error)
}

type Command interface { // GORM: Create, Save, Update, Updates, FirstOrCreate, FirstOrInit, ...
	Create(ctx context.Context, params CreateParams) (*Entity, error)
	Update(ctx context.Context, entity *Entity) error
}
