package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/jwt"
	"github.com/mzfarshad/music_store_api/pkg/logger"
)

const AuthenticateUserKey = "Authenticate"

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		log := logger.GetLogger(ctx)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		tokenUser, err := jwt.ValidateToken(ctx, authHeader)
		if err != nil {
			if customErr, ok := err.(*apperr.CustomErr); ok {
				customErr.Code = apperr.StatusUnauthorized
				log.Error(ctx, "", customErr)
				c.IndentedJSON(customErr.Code, gin.H{"Message": customErr.Message})
				c.Abort()
				return
			}
		}
		c.Set(AuthenticateUserKey, tokenUser)
		c.Next()
	}
}
