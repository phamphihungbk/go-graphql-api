package app

import (
	"net/http"

	"github.com/phamphihungbk/go-graphql-api/internal/config"
	"github.com/phamphihungbk/go-graphql-api/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql-api/internal/graphql/generated"
	"github.com/phamphihungbk/go-graphql-api/internal/resolver"
)

type Router struct {
	isBooted     bool
	resolverRoot *resolver.Resolver
	routerRoot   *gin.Engine
	authService  *service.AuthenticationService
	logger       *service.Logger
	config       config.Server
}

type IRouter interface {
	Boot()
}

func NewRoute(
	resolver *resolver.Resolver,
	authService *service.AuthenticationService,
	logger *service.Logger,
	config config.Server,
) *Router {
	router := gin.Default()

	return &Router{
		isBooted:     false,
		resolverRoot: resolver,
		routerRoot:   router,
		authService:  authService,
		logger:       logger,
		config:       config,
	}
}

func (r *Router) Boot() {
	if r.isBooted {
		return
	}

	port := r.config.Port

	r.setupRoutes()
	r.setupMiddlewares()
	_ = r.routerRoot.Run(port)

	r.isBooted = true
}

func (r *Router) setupMiddlewares() {
	r.routerRoot.Use(AuthenticateUser(r.authService))
	r.routerRoot.Use(ConvertGinContextToContext())
	r.routerRoot.Use(HandleException(r.logger))
}

func (r *Router) setupRoutes() {
	r.routerRoot.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.routerRoot.POST("/query", r.graphqlHandler())
	r.routerRoot.GET("/", r.playgroundHandler("/query"))
}

func (r *Router) graphqlHandler() gin.HandlerFunc {
	conf := generateGraphQLConfig(r.resolverRoot)
	exec := generated.NewExecutableSchema(conf)
	h := handler.NewDefaultServer(exec)

	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

func generateGraphQLConfig(resolvers *resolver.Resolver) generated.Config {
	return generated.Config{
		Resolvers: resolvers,
	}
}

func (r *Router) playgroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL", path)

	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
