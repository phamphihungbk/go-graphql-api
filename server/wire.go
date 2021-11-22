// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeApp(dbCfg string) (Application, error) {
	panic(wire.Build(ProvideControllers, ProvideServices, ProvideRepositories, ProvideDBConnection))

	return Application{}, nil
}
