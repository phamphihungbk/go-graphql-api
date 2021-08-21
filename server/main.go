package main

import (
	"github.com/phamphihungbk/go-graphql/app/container"
)

func main() {
	err := container.initApp()

	if err != nil {
		panic(err)
	}
}
