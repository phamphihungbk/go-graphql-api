package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(builder *di.Builder) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	userController := dic.Container.Get(dic.UserController).(*controller.UserController)

	// Ping
	//
	// Get Ping and reply Pong
	//
	//     Responses:
	//       200:
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// User
	//
	// Get user data
	//
	//     Responses:
	//       200: UserResponse
	r.GET("/users/:id", userController.Get)

	// New user
	//
	// Create new user
	//
	//     Responses:
	//       200: UserResponse
	r.POST("/users", userController.Create)

	// Update user
	//
	// Update existing user
	//
	//     Responses:
	//       200: UserResponse
	r.PUT("/users/:id", userController.Update)

	// Delete user
	//
	// Delete existing user
	//
	//     Responses:
	//       200:
	r.DELETE("/users/:id", userController.Delete)

	return r
}
