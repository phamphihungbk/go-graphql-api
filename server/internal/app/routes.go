package app

import (
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql/internal/resolver"
	"github.com/99designs/gqlgen/handler"
	"net/http"
	"github.com/phamphihungbk/go-graphql/graphql/generated"
)

func NewRoute(controller *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.Use(GinContextToContextMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	return r
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.NewUserResolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
