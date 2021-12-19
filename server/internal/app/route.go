package app

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql/internal/graphql/generated"
	"github.com/phamphihungbk/go-graphql/internal/resolver"
	"github.com/phamphihungbk/go-graphql/internal/service"
)

func NewRoute(userService service.IUserService) *gin.Engine {
	r := gin.Default()
	r.Use(GinContextToContextMiddleware())
	r.Use(LoggerToFile())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/query", graphqlHandler(userService))
	r.GET("/", playgroundHandler("/query"))

	return r
}

func graphqlHandler(userService service.IUserService) gin.HandlerFunc {
	conf := generated.Config{
		Resolvers: resolver.NewResolver(
			userService,
		),
	}
	exec := generated.NewExecutableSchema(conf)
	h := handler.NewDefaultServer(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

func playgroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", path)

	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
