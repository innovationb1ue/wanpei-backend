package controller

import "go.uber.org/fx"

func RegisterControllers() fx.Option {
	return fx.Module("controllers", fx.Invoke(
		ping,
		UserRoutes,
		GameRoutes,
		MatchRoutes,
	))
}
