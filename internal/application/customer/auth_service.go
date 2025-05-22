package customer

import (
	"context"
	"errors"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/errs"
)

func NewAuthService(userRepo user.Repository) auth.CustomerUseCase {
	return &authService{userRepo: userRepo}
}

type authService struct {
	userRepo user.Repository
}

func (s *authService) SignIn(ctx context.Context, email, password string) (*auth.PairToken, error) {
	customer, err := s.userRepo.First(ctx, user.Where{
		Email: email,
		Type:  user.TypeCustomer,
	})
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errs.New(errs.NotFound, "customer not found")
	}
	if !customer.Active {
		return nil, errs.New(errs.Unauthorized, "your account has been deactivated")
	}
	if err = customer.CompareHashAndPassword(password); err != nil {
		return nil, err
	}
	access, err := auth.NewAccessToken(customer.Email, customer.Type, customer.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{
		Access: *access,
	}, nil
}

func (s *authService) Signup(ctx context.Context, name, email, password string) (*auth.PairToken, error) {
	usrByEmail, err := s.userRepo.First(ctx, user.Where{Email: email})
	if err != nil && !errors.Is(err, errs.NotFound.Err()) {
		return nil, err
	}
	if usrByEmail != nil { // user exists
		return nil, errs.New(errs.Duplication, "email already taken")
	}
	usr, err := s.userRepo.Create(ctx, user.CreateParams{
		Name:     name,
		Email:    email,
		Password: password,
		Type:     user.TypeCustomer,
	})
	if err != nil {
		return nil, err
	}
	access, err := auth.NewAccessToken(usr.Email, usr.Type, usr.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{Access: *access}, nil
}
