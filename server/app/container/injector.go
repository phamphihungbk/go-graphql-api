package container

import (
	"github.com/google/wire"
)

func initializeEnv() {
	wire.Build()
}

func initializeDbConnection() {
	wire.Build()
}

func initializeRepositories() {
	wire.Build()
}

func initializeServices() {
	wire.Build()
}

func initializeControllers() {
	wire.Build()
}

func initializeRoutes() {
	wire.Build()
}
