package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewLogger),
	fx.Provide(NewGinRouter),
)
