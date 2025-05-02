package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mzfarshad/music_store_api/config"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/logger"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Email    string
	UserType user.Type
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
		UserType: userType,
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

func ValidateToken(ctx context.Context, tkn string) (*UserClaims, error) {
	log := logger.GetLogger(ctx)
	tokenUser := &UserClaims{}
	token, err := jwt.Parse(tkn, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Get().Jwt.Access.Secret), nil
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
		tokenUser.UserType = (claim["user_type"]).(user.Type)
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
