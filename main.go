package main

import (
	"encoding/gob"
	"go.uber.org/fx"
	"wanpei-backend/controller"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/repo"
	"wanpei-backend/server"
	"wanpei-backend/services"
	"wanpei-backend/worker"
)

func main() {
	// register models for encoding
	gob.Register(models.User{})
	gob.Register(models.UserInsensitive{})
	// start app
	app := fx.New(
		// get environment variables
		fx.Provide(server.GetEnv),
		// provide infrastructures constructors
		fx.Provide(mapper.NewDbConn, server.NewApp, models.NewSessionStore, server.NewSettings),
		// creat local repos
		repo.CreateRepo(),
		// register all mappers
		mapper.RegisterMapper(),
		// register services
		services.RegisterServices(),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		// Start matchmaking worker goroutine
		fx.Provide(worker.NewMatch),
		fx.Invoke(worker.MatchWorker),
		// run gin app
		fx.Invoke(server.Run),
	)
	app.Run() // run forever
}
