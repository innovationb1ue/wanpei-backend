package main

import (
	"encoding/gob"
	"go.uber.org/fx"
	"wanpei-backend/controller"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/server"
	"wanpei-backend/services"
)

func main() {
	// register user type to be saved in sessions
	// todo: move it to a independent file
	gob.Register(models.User{})
	// start app
	app := fx.New(
		// provide infrastructures constructors
		fx.Provide(models.NewDbConn, server.NewApp, models.NewSessionStore, server.NewSettings,
			models.NewRedisConn),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		mapper.RegisterMapper(),
		services.RegisterServices(),
		// invoke functions should run before the app start
		fx.Invoke(server.Run))
	app.Run() // run forever
}
