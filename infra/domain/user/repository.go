package user

import "context"

type Repository interface {
	Query
	Command
}

// TODO: @Farshad Search about CQRS

type Query interface { // GORM: First, Last, Take, Find
	First(ctx context.Context, email string) (*Entity, error)
}

type Command interface { // GORM: Create, Save, Update, Updates, FirstOrCreate, FirstOrInit, ...
	Create(ctx context.Context, params CreateParams) (*Entity, error)
}
