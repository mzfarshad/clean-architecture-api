package application

import (
	"context"

	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewCustomerService(userRepo user.Repository) user.CustomerUseCase {
	return &customerService{
		userRepo: userRepo,
	}
}

type customerService struct {
	userRepo user.Repository
}

func (s *customerService) UpdateMyName(ctx context.Context, name, email string) error {
	usr, err := s.userRepo.FirstByEmail(ctx, email)
	if err != nil {
		return err
	}
	usr.Name = name
	if err := s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}
