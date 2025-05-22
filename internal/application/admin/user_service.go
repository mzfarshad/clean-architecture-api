package admin

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewUserService(userRepo user.Repository) user.AdminUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo user.Repository
}

func (s *userService) DeactivateUser(ctx context.Context, userId uint, reason string) error {
	usr, err := s.userRepo.First(ctx, user.Where{Id: userId})
	if err != nil {
		return err
	}
	if !usr.Active {
		return fmt.Errorf("user with id %d is deactive, reason: %s", userId, usr.InactiveReason)
	}
	usr.Active = false
	usr.InactiveReason = reason
	if err := s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}

func (s *userService) ReactivateUser(ctx context.Context, userId uint) error {
	usr, err := s.userRepo.First(ctx, user.Where{Id: userId})
	if err != nil {
		return err
	}
	if usr.Active {
		return fmt.Errorf("user has been active")
	}
	usr.Active = true
	if err := s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}

func (s *userService) SearchInUsers(ctx context.Context,
	params user.SearchParams) (*user.PaginationParams, error) {
	usrPagination, err := s.userRepo.Find(ctx, params)
	if err != nil {
		return nil, err
	}
	return usrPagination, nil
}

func (s *userService) UpdateMyProfile(ctx context.Context, name string) error {
	claims, err := auth.MustClaimed(ctx, user.TypeAdmin)
	if err != nil {
		return err
	}
	usr, err := s.userRepo.First(ctx, user.Where{
		Id:   claims.ID,
		Type: claims.UserType,
	})
	if err != nil {
		return err
	}
	usr.Name = name
	if err = s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}
