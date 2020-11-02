package models

import (
	"golang-fx/infrastructure"

	"go.uber.org/fx"
)

// Module exported from models package
var Module = fx.Options(
	fx.Provide(NewMigrations),
)

// Migrations migration data type
type Migrations struct {
	db     infrastructure.Database
	logger infrastructure.Logger
	models []ModelList
}

// ModelList interface
type ModelList interface {
}

// NewMigrations creates new migrations instance
func NewMigrations(db infrastructure.Database, logger infrastructure.Logger) Migrations {
	return Migrations{
		db:     db,
		logger: logger,
		models: []ModelList{
			User{}, Product{},
		},
	}
}

// Migrate function to call when migrating
func (m Migrations) Migrate() {
	m.logger.Zap.Info("Automigrating schemas.")
	for _, model := range m.models {
		err := m.db.DB.AutoMigrate(model)
		if err != nil {
			m.logger.Zap.Fatalf("Migration failure: ", err)
		}
	}
}
