package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mzfarshad/music_store_api/config"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Email    string
	UserType string
	ID       uint
}

func NewAccessToken(email string, userType user.Type, id uint) (*Token, error) {
	cfg := config.Get()
	now := time.Now()
	expiresAt := now.Add(cfg.Jwt.Access.TTL)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Email:    email,
		UserType: string(userType),
		ID:       id,
	})
	secretKey := []byte(cfg.Jwt.Access.Secret)
	tokenStr, err := jwtToken.SignedString(secretKey)
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed make token string",
			apperr.TypeInternal,
			err.Error(),
		)
		return nil, customErr
	}
	return &Token{
		Raw:       tokenStr,
		ExpiresAt: time.Now().Add(cfg.Jwt.Access.TTL),
	}, nil
}
