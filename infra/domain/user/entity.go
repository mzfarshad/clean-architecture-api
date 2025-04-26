package user

import (
	"github.com/mzfarshad/music_store_api/infra/domain"
	"golang.org/x/crypto/bcrypt"
)

type Entity struct {
	domain.Entity
	Name         string
	Email        string
	passwordHash string
	Type         Type
}

func (e Entity) CompareHashAndPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.passwordHash), []byte(password))
}
