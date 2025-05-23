package customer

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain"
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

func (s *userService) UpdateMyName(ctx context.Context, name string) (*user.Entity, error) {
	claims, err := auth.MustClaimed(ctx, domain.Customer)
	if err != nil {
		return nil, err
	}
	customer, err := s.userRepo.First(ctx, user.Where{
		Id:   claims.ID,
		Type: domain.Customer,
	})
	if err != nil {
		return nil, err
	}
	updateParams := user.UpdateParams{Name: name}
	updateParams.Where.Id = customer.Id
	return s.userRepo.Update(ctx, updateParams)
}
