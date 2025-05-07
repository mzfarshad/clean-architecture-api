package admin

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

type CliService interface {
	CreateAdmin(email, name, pass string) error
}

func NewCliService(userRepo user.Repository) CliService {
	return &adminService{
		repo: userRepo,
	}
}

type adminService struct {
	repo user.Repository
}

func (s *adminService) CreateAdmin(email, name, pass string) error {
	var admin user.CreateParams
	admin.Name = name
	admin.Email = email
	admin.Password = pass
	admin.Type = user.TypeAdmin
	_, err := s.repo.Create(context.Background(), admin)
	return err
}
