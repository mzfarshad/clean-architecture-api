package user

import (
	"github.com/mzfarshad/music_store_api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Entity struct {
	domain.Entity
	Name           string
	Email          string
	PasswordHash   string
	InactiveReason string
	Type           Type
	Status         bool
}

func (e Entity) CompareHashAndPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.PasswordHash), []byte(password))
}
