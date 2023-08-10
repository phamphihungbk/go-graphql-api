package app

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/phamphihungbk/go-graphql-api/internal/service"

	"github.com/gin-gonic/gin"
)

type ContextKey string

func ConvertGinContextToContext() gin.HandlerFunc {
	key := ContextKey("GinContextKey")
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), key, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func HandleException(logger *service.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURL := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Info("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURL,
		)
	}
}

func AuthenticateUser(authService *service.AuthenticationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := service.ValidateToken(token)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		email, _ := claims["email"].(string)
		userId, _ := claims["sub"].(string)

		if !authService.IsValidUser(email) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "UserID", userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
