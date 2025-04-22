package authuser

import (
	"context"

	"github.com/mzfarshad/music_store_api/internal/api/presenter"
	"github.com/mzfarshad/music_store_api/internal/models"
	userrepo "github.com/mzfarshad/music_store_api/internal/repo/user_repo"
)

type authUserService struct {
	Repo userrepo.AuthUserRepository
}

func NewAuthUserService() *authUserService {
	db, _ := models.NewPostgresConnection()
	repo := userrepo.NewAuthUserRepo(db)
	return &authUserService{Repo: repo}
}

func (a *authUserService) FindEmail(ctx context.Context, email string) (models.User, error) {
	return a.Repo.FindEmail(ctx, email)
}

func (a *authUserService) SaveUser(ctx context.Context, user presenter.SignUpUser) error {
	return a.Repo.SaveUser(ctx, user)
}
