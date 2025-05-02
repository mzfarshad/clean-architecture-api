package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mzfarshad/music_store_api/conf"
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
	jwtAccessTokenTTL := 24 * time.Hour // // TODO: add TTL to jwt config
	now := time.Now()
	expiresAt := now.Add(jwtAccessTokenTTL)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Email:    email,
		UserType: string(userType),
		ID:       id,
	})
	secretKey := []byte(conf.Get().JWT().SecretKey)
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
		ExpiresAt: time.Now().Add(jwtAccessTokenTTL),
	}, nil
}
