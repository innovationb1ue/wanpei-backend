package main

import (
	"go.uber.org/fx"
	"wanpei-backend/controller"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
	"wanpei-backend/server"
	"wanpei-backend/services"
)

func main() {

	app := fx.New(
		// provide infrastructures constructors
		fx.Provide(models.NewDbConn, server.NewApp, models.NewSessionMgr, NewSettings),
		// register all controllers by providing fx.Option
		controller.RegisterControllers(),
		mapper.RegisterMapper(),
		services.RegisterServices(),
		// invoke functions should run before the app start
		fx.Invoke(server.Run))
	app.Run() // run forever
}
