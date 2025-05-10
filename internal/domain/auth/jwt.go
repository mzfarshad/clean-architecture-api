package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mzfarshad/music_store_api/config"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"time"
)

var (
	ErrTokenInvalid            = errors.New("token is invalid")
	ErrTokenExpired            = errors.New("token is expired")
	ErrTokenMalformed          = errors.New("token is malformed")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
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
	tokenClaims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tkn, tokenClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(config.Get().Jwt.Access.Secret), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrTokenMalformed
		default:
			return nil, fmt.Errorf("token parse error: %w", err)
		}
	}
	if !token.Valid {
		return nil, ErrTokenInvalid
	}
	return tokenClaims, nil
}
