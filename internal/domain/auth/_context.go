package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/mzfarshad/music_store_api/internal/domain/user"
)

type ctxKey string

const (
	CtxKeyAuthUser ctxKey = "Auth-User"
)

func MustClaimed(ctx context.Context, roles ...user.Type) (*UserClaims, error) {
	claims, ok := ctx.Value(CtxKeyAuthUser).(*UserClaims)
	if !ok {
		return nil, errors.New("you're not authenticated.") // errs.Unauthorized
	}
	if roles != nil && len(roles) > 0 {
		if !claims.UserType.Is(roles[0], roles[1:]...) {
			return nil, errs.New(errs.Forbidden, fmt.Sprintf("access denied for a %s", claims.UserType))
		}
	}
	return claims, nil
}
