package user

import "context"

// AdminUseCase represents what an TypeAdmin can do with user Entity.
type AdminUseCase interface {
	ReactivateUser(ctx context.Context, userId uint) error
	DeactivateUser(ctx context.Context, userId uint, reason string) error
	SearchInUsers(ctx context.Context, params SearchParams) (*PaginationParams, error)
	UpdateMyProfile(ctx context.Context, name, email string) error // added email to parmas for searching user
}

// CustomerUseCase represents what a TypeCustomer can do with user Entity.
type CustomerUseCase interface {
	UpdateMyName(ctx context.Context, name, email string) error // added email to parmas for searching user
}

type CliUseCase interface {
	IncreaseUsersCredit(ctx context.Context, amount uint) error
}
