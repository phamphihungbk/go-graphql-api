package container

import (
	"go.uber.org/dig"
	"log"
)

var Container *dig.Container

func DatabaseConnection(config *Config) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.GetConfig()

	db, err := gorm.Open(postgres.Open(config.Postgres.GetConnectionInfo()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func Controllers() {
	return
	{
	}

}

func Models() {
	return
	{
	}
}

func Environment() {
	return
	{
	}
}

func NewServer() *Server {
	return &Server{
	}
}

func BuildContainer() *dig.Container {
	Container := dig.New()
	Container.Provide(Environment)
	Container.Provide(DatabaseConnection)
	Container.Provide(Models)
	Container.Provide(Controllers)
	Container.Provide(NewServer)

	return Container
}
