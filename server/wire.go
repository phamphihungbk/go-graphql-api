//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql/app"
	"github.com/phamphihungbk/go-graphql/app/controllers"
	"github.com/phamphihungbk/go-graphql/configs"
	"github.com/phamphihungbk/go-graphql/src/repositories"
	"github.com/phamphihungbk/go-graphql/src/services"
)

func InitializeApp(appCfg *configs.AppCfg, dbCfg *configs.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
		app.NewRouter,
		NewApplication,
	)

	return &Application{}, nil
}
