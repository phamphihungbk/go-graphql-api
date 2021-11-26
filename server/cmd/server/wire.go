//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql/internal/app"
	"github.com/phamphihungbk/go-graphql/internal/repository"
	"github.com/phamphihungbk/go-graphql/internal/service"
)

func InitializeApp(appCfg *config.AppCfg, dbCfg *config.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repository.NewUserRepository,
		service.NewUserService,
		app.NewRoute,
		NewApplication,
	)

	return &Application{}, nil
}
