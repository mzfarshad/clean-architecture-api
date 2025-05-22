package customer

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewUserService(userRepo user.Repository) user.CustomerUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo user.Repository
}

func (s *userService) UpdateMyName(ctx context.Context, name string) error {
	claims, err := auth.MustClaimed(ctx, user.TypeCustomer)
	if err != nil {
		return err
	}
	customer, err := s.userRepo.First(ctx, user.Where{
		Id:   claims.ID,
		Type: user.TypeCustomer,
	})
	if err != nil {
		return err
	}
	customer.Name = name
	if err = s.userRepo.Update(ctx, customer); err != nil {
		return err
	}
	return nil
}
