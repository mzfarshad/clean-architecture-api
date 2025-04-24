package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mzfarshad/music_store_api/pkg/logger"
)

const (
	requestIDKey    = "RequestID"
	requestIdHeader = "X-Request-ID"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		c.Set(requestIDKey, reqID)
		c.Writer.Header().Set(requestIdHeader, reqID)

		log := logger.Get("DEBUG")
		log.Conf(func(opt *logger.Option) {
			opt.RequestID = reqID
			opt.ShowFileName = true
			opt.ShowFunctionName = true
			opt.ShowLine = true
			opt.ShowTimeStamp = true
		})

		ctx := logger.WithLogger(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
