//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql-api/internal/app"
	"github.com/phamphihungbk/go-graphql-api/internal/config"
	"github.com/phamphihungbk/go-graphql-api/internal/repository"
	"github.com/phamphihungbk/go-graphql-api/internal/resolver"
	"github.com/phamphihungbk/go-graphql-api/internal/service"
)

func InitializeApp(config config.Server) (*App, error) {
	wire.Build(
		app.NewDatabaseConnection,
		repository.NewUserRepository,
		service.NewUserService,
		resolver.NewResolver,
		service.NewAuthenticationService,
		app.NewRoute,
		NewApp,
	)

	return &App{}, nil
}
