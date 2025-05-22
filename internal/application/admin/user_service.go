package admin

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/pkg/search"
)

func NewUserService(userRepo user.Repository) user.AdminUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo user.Repository
}

func (s *userService) DeactivateUser(ctx context.Context, userId uint, reason string) (*user.Entity, error) {
	usr, err := s.userRepo.First(ctx, user.Where{Id: userId})
	if err != nil {
		return nil, err
	}
	if !usr.Active {
		return nil, errs.New(errs.Unprocessable, fmt.Sprintf("user #%d has been already deactived because: %s", userId, usr.InactiveReason))
	}
	updateParams := user.UpdateParams{
		InactiveReason: reason,
		Active:         dto.NewOptional(false),
	}
	updateParams.Where.Id = userId
	return s.userRepo.Update(ctx, updateParams)
}

func (s *userService) ReactivateUser(ctx context.Context, userId uint) (*user.Entity, error) {
	usr, err := s.userRepo.First(ctx, user.Where{Id: userId})
	if err != nil {
		return nil, err
	}
	if usr.Active {
		return usr, nil
	}
	updateParams := user.UpdateParams{
		Active: dto.NewOptional(true),
	}
	updateParams.Where.Id = userId
	return s.userRepo.Update(ctx, updateParams)
}

func (s *userService) SearchInUsers(ctx context.Context, p *search.Pagination[user.SearchParams]) ([]*user.Entity, error) {
	return s.userRepo.Search(ctx, p)
}

func (s *userService) UpdateMyProfile(ctx context.Context, name string) (*user.Entity, error) {
	claims, err := auth.MustClaimed(ctx, user.TypeAdmin)
	if err != nil {
		return nil, err
	}
	usr, err := s.userRepo.First(ctx, user.Where{
		Id:   claims.ID,
		Type: claims.UserType,
	})
	if err != nil {
		return nil, err
	}
	updateParams := user.UpdateParams{Name: name}
	updateParams.Where.Id = usr.Id
	return s.userRepo.Update(ctx, updateParams)
}
