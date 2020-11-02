package bootstrap

import (
	"context"
	"golang-fx/api/controllers"
	"golang-fx/api/models"
	"golang-fx/api/repository"
	"golang-fx/api/routes"
	"golang-fx/infrastructure"
	"os"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	repository.Module,
	routes.Module,
	infrastructure.Module,
	models.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.GinRouter,
	routes routes.Routes,
	logger infrastructure.Logger,
	database infrastructure.Database,
	migrations models.Migrations,

) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-----------------------------")
			logger.Zap.Info("------- GO-FX-API 🚀 -------")
			logger.Zap.Info("-----------------------------")

			go func() {
				migrations.Migrate()
				routes.Setup()
				handler.Gin.Run(os.Getenv("SERVER_PORT"))
			}()
			return nil

		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping CRON Jobs")

			logger.Zap.Info("Stopping Application")
			sqlDB, _ := database.DB.DB()
			sqlDB.Close()
			return nil
		},
	})
}
