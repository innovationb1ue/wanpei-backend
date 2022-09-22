package repo

import "go.uber.org/fx"

func CreateRepo() fx.Option {
	return fx.Module("repo", fx.Provide(
		NewHub,
	))
}
