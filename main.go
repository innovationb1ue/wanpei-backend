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
	gob.Register(models.User{})
	gob.Register(models.UserInsensitive{})
	// start app
	app := fx.New(
		// provide infrastructures constructors
		fx.Provide(mapper.NewDbConn, server.NewApp, models.NewSessionStore, server.NewSettings),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		mapper.RegisterMapper(),
		services.RegisterServices(),
		repo.CreateRepo(),
		fx.Provide(worker.NewMatch),
		fx.Invoke(worker.MatchWorker),
		// invoke functions should run before the app start
		fx.Invoke(server.Run),
	)
	app.Run() // run forever
}
