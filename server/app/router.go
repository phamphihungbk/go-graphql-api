package app

import (
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql/app/controllers"
	"net/http"
)

func NewRouter(controller *controllers.UserController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/users/:id", controller.Get)
	r.POST("/users", controller.Create)
	r.PUT("/users/:id", controller.Update)
	r.DELETE("/users/:id", controller.Delete)

	return r
}
