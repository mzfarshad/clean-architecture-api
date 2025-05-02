package customer

import (
	"context"

	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewAuthService(userRepo user.Repository) auth.CustomerUseCase {
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
	access, err := auth.NewAccessToken(usr.Email, usr.Type, usr.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{
		Access: *access,
	}, nil
}

func (s authService) Signup(ctx context.Context, name, email, password string) (*auth.PairToken, error) {
	var customer user.CreateParams
	_, err := s.userRepo.FirstByEmail(ctx, email)
	if err == nil {
		return nil, err
	}
	customer.Email = email
	customer.Name = name
	customer.Password = password
	customer.Type = user.TypeCustomer

	usr, err := s.userRepo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	access, err := auth.NewAccessToken(usr.Email, usr.Type, usr.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{Access: *access}, nil
}
