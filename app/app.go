package app

import (
	"context"
	"golang-fx/api/controller"
	"golang-fx/api/repository"
	"golang-fx/api/route"
	"golang-fx/api/service"
	"golang-fx/infrastructure"
	"golang-fx/model"
	"os"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controller.Module,
	repository.Module,
	service.Module,
	route.Module,
	infrastructure.Module,
	model.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	router infrastructure.GinRouter,
	routes route.Routes,
	logger infrastructure.Logger,
	database infrastructure.Database,
	migrations model.Migrations,

) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-----------------------------")
			logger.Zap.Info("------- golang-fx 🚀 -------")
			logger.Zap.Info("-----------------------------")
			migrations.Migrate()
			routes.Setup()
			go func() {
				router.Gin.Run(os.Getenv("SERVER_PORT"))
			}()

			return nil

		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			sqlDB, _ := database.DB.DB()
			sqlDB.Close()
			return nil
		},
	})
}
