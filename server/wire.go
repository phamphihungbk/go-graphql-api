//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/phamphihungbk/go-graphql/app/controllers"
	"github.com/phamphihungbk/go-graphql/configs"
	"github.com/phamphihungbk/go-graphql/database"
	"github.com/phamphihungbk/go-graphql/src/repositories"
	"github.com/phamphihungbk/go-graphql/src/services"
	"github.com/phamphihungbk/go-graphql/app"
)

func InitializeApp(dbCfg *configs.DBCfg) (*Application, error){
	wire.Build(
		database.NewDBConnection,
		repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
		app.NewRouter,
		NewApplication,
	)

	return nil, nil
}
