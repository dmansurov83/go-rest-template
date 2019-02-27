package main

import (
	"./db"
	"./env"
	"./services"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(env.NewConfig)
	container.Provide(db.CreateDB)
	container.Provide(db.NewPersonRepository)
	container.Provide(services.NewPersonService)
	container.Provide(NewServer)
	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
