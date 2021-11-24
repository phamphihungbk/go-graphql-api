package app

import (
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql/app/controller"
	"net/http"
)

func NewRoute(controller *controller.UserController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
	//r.GET("/users/:id", controller.Get)
	//r.POST("/users", controller.Create)
	//r.PUT("/users/:id", controller.Update)
	//r.DELETE("/users/:id", controller.Delete)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
