package jwt

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/conf"

	"github.com/golang-jwt/jwt/v5"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/logger"
)

type TokenUser struct {
	Email    string
	UserType string
	ID       uint
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

func ValidateToken(ctx context.Context, tkn string) (*TokenUser, error) {
	log := logger.GetLogger(ctx)
	tokenUser := &TokenUser{}
	token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(conf.Get().JWT().SecretKey), nil
	})
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed validate token",
			apperr.TypeApi,
			err.Error(),
		)
		log.Error(ctx, "", customErr)
		return nil, customErr
	}
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenUser.Email = (claim["email"]).(string)
		tokenUser.UserType = (claim["user_type"]).(string)
		userID := (claim["id"]).(float64)
		tokenUser.ID = uint(userID)
		return tokenUser, nil
	}
	customErr := apperr.NewAppErr(
		apperr.StatusInternalServerError,
		"invalidate token",
		apperr.TypeApi,
		"",
	)
	return nil, customErr
}
