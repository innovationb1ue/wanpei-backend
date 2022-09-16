package main

import (
	"go.uber.org/fx"
	"wanpei-backend/controller"
	"wanpei-backend/models"
	"wanpei-backend/server"
)

func main() {

	app := fx.New(
		// provide infrastructures constructors
		fx.Provide(models.NewDbConn, server.NewApp, models.NewSessionMgr, NewSettings),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		// invoke functions should run before the app start
		fx.Invoke(server.Run))
	app.Run() // run forever
}
