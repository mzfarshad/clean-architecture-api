package presenter

import (
	"github.com/mzfarshad/music_store_api/internal/domain/auth"
	"github.com/mzfarshad/music_store_api/rest"
	"time"
)

type AuthToken struct {
	rest.DTO  `json:"-"`
	Raw       string    `json:"raw"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewAuthToken(token auth.Token) *AuthToken {
	return &AuthToken{
		Raw:       token.Raw,
		ExpiresAt: token.ExpiresAt,
	}
}
