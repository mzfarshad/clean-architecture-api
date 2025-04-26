package auth

import (
	"context"
)

type UseCase interface {
	SignIn(ctx context.Context, email, password string) (*PairToken, error)
	Signup(ctx context.Context, email, password string) (*PairToken, error)
}
