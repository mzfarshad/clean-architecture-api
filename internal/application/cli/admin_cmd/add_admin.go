package admin_cmd

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

type Service interface {
	CreateAdmin(email, name, pass string) error
}

func NewAdminCmdService(userRepo user.Repository) Service {
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
