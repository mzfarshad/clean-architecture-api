package application

import (
	"context"

	"github.com/mzfarshad/music_store_api/infra/domain/auth"
	"github.com/mzfarshad/music_store_api/infra/domain/user"
)

func NewAuthService(userRepo user.Repository) auth.UseCase {
	return &authService{userRepo: userRepo}
}

type authService struct {
	userRepo user.Repository
}

func (s authService) SignIn(ctx context.Context, email, password string) (*auth.PairToken, error) {
	usr, err := s.userRepo.FirstByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err = usr.CompareHashAndPassword(password); err != nil {
		return nil, err
	}
	access, err := auth.NewAccessToken(usr.Email, string(usr.Type), usr.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{
		Access: *access,
	}, nil
}

func (s authService) Signup(ctx context.Context, email, password string) (*auth.PairToken, error) {
	//TODO implement me
	panic("implement me")
}
