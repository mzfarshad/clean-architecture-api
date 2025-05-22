package auth

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/errs"
)

type ctxKey string

const (
	CtxKeyAuthUser ctxKey = "Auth-User"
)

func MustClaimed(ctx context.Context, roles ...user.Type) (*UserClaims, error) {
	claims, exists := ctx.Value(CtxKeyAuthUser).(*UserClaims)
	if !exists {
		return nil, errs.New(errs.Unauthorized, "you're not authenticated")
	}
	if roles != nil && len(roles) > 0 {
		if !claims.UserType.Is(roles[0], roles[1:]...) {
			return nil, errs.New(errs.Forbidden, fmt.Sprintf("access denied for a %s", claims.UserType))
		}
	}
	return claims, nil
}
