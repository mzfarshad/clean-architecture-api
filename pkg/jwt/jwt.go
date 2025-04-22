package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mzfarshad/music_store_api/internal/conf"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type TokenUser struct {
	Email    string
	UserType string
	id       uint
}

func NewAccessToken(email, userType string, id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"user_type": userType,
		"id":        id,
	})
	secretKey := []byte(conf.Get().JWT().SecretKey)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed make token string",
			apperr.TypeInternal,
			err.Error(),
		)
		return "", customErr
	}
	return tokenStr, nil
}
