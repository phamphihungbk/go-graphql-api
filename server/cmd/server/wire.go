//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql/config"
	"github.com/phamphihungbk/go-graphql/internal/app"
	"github.com/phamphihungbk/go-graphql/internal/repository"
	"github.com/phamphihungbk/go-graphql/internal/resolver"
	"github.com/phamphihungbk/go-graphql/internal/service"
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
