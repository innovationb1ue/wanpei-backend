package services

import "go.uber.org/fx"

func RegisterServices() fx.Option {
	return fx.Module("services", fx.Provide(
		NewUser,
		NewGame,
		NewMatch,
		NewToken,
		NewSocket,
		NewHub,
	))
}
