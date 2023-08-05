//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql-api/config"
	"github.com/phamphihungbk/go-graphql-api/internal/app"
	"github.com/phamphihungbk/go-graphql-api/internal/repository"
	"github.com/phamphihungbk/go-graphql-api/internal/resolver"
	"github.com/phamphihungbk/go-graphql-api/internal/service"
)

func InitializeApp(dbCfg *config.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repository.NewUserRepository,
		service.NewUserService,
		resolver.NewResolver,
		app.NewRoute,
		NewApplication,
	)

	return &Application{}, nil
}
