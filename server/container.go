// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeApp(connectionInfo string) error {
	wire.Build(ProvideConnections)

	return nil
}
