package app

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ContextKey string

func GinContextToContextMiddleware() gin.HandlerFunc {
	key := ContextKey("GinContextKey")
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), key, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
