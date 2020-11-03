package model

import (
	"go.uber.org/fx"
)

// Module exported from models package
var Module = fx.Options(
	fx.Provide(NewMigrations),
)
