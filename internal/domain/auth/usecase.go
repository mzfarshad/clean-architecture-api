package auth

import (
	"context"
)

type CustomerUseCase interface {
	SignIn(ctx context.Context, email, password string) (*PairToken, error)
	Signup(ctx context.Context, name, email, password string) (*PairToken, error)
}

type AdminUseCase interface {
	SingIn(ctx context.Context, email, password string) (*PairToken, error)
}
