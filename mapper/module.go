package mapper

import "go.uber.org/fx"

func RegisterMapper() fx.Option {
	return fx.Module("mapper", fx.Provide(
		NewUser,
		NewGame,
	))
}
