package admin

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

func NewAuthService(repo user.Repository) auth.AdminUseCase {
	return &authService{
		repo: repo,
	}
}

type authService struct {
	repo user.Repository
}

func (s *authService) SingIn(ctx context.Context, email, password string) (*auth.PairToken, error) {
	admin, err := s.repo.First(ctx, user.Where{
		Email: email,
		Type:  user.TypeAdmin,
	})
	if err != nil {
		return nil, err
	}
	if err = admin.CompareHashAndPassword(password); err != nil {
		return nil, err
	}
	access, err := auth.NewAccessToken(admin.Email, admin.Type, admin.Id)
	if err != nil {
		return nil, err
	}
	return &auth.PairToken{
		Access: *access,
	}, nil
}
