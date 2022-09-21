package main

import (
	"encoding/gob"
	"go.uber.org/fx"
	"wanpei-backend/controller"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/server"
	"wanpei-backend/services"
	"wanpei-backend/worker"
)

func main() {
	// register user type to be saved in sessions
	// todo: move it to a independent file
	gob.Register(models.User{})
	// start app
	app := fx.New(
		// provide infrastructures constructors
		fx.Provide(mapper.NewDbConn, server.NewApp, models.NewSessionStore, server.NewSettings),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		mapper.RegisterMapper(),
		services.RegisterServices(),
		fx.Provide(worker.NewMatch),
		fx.Invoke(worker.MatchWorker),
		// invoke functions should run before the app start
		fx.Invoke(server.Run),
	)
	app.Run() // run forever
}
