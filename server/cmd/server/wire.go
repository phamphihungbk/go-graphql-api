//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql/app"
	"github.com/phamphihungbk/go-graphql/app/controller"
	"github.com/phamphihungbk/go-graphql/src/repository"
	"github.com/phamphihungbk/go-graphql/src/service"
)

func InitializeApp(appCfg *config.AppCfg, dbCfg *config.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		app.NewRoute,
		NewApplication,
	)

	return &Application{}, nil
}
