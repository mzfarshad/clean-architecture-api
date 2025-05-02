package admin

import (
	"context"
	"fmt"

	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewUserService(userRepo user.Repository) user.AdminUseCase {
	return &adminService{
		userRepo: userRepo,
	}
}

type adminService struct {
	userRepo user.Repository
}

func (s *adminService) DeactivateUser(ctx context.Context, userId uint, reason string) error {
	usr, err := s.userRepo.FirstById(ctx, userId)
	if err != nil {
		return err
	}
	if !usr.Status {
		return fmt.Errorf("user with id %d is deactive, reason: %s", userId, usr.InactiveReason)
	}
	usr.Status = false
	usr.InactiveReason = reason
	if err := s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}

func (s *adminService) ReactivateUser(ctx context.Context, userId uint) error {
	usr, err := s.userRepo.FirstById(ctx, userId)
	if err != nil {
		return err
	}
	if usr.Status {
		return fmt.Errorf("user has been active")
	}
	usr.Status = true
	if err := s.userRepo.Update(ctx, usr); err != nil {
		return err
	}
	return nil
}

func (s *adminService) SearchInUsers(ctx context.Context,
	params user.SearchParams) (*user.PaginationParams, error) {
	if params.Limit < 1 {
		params.Limit = 10
	}
	if params.Page < 1 {
		params.Page = 1
	}
	usrPagination, err := s.userRepo.Find(ctx, params)
	if err != nil {
		return nil, err
	}
	return usrPagination, nil
}

func (s *adminService) UpdateMyProfile(ctx context.Context, name, email string) error {
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
