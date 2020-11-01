package routes

import (
	"golang-fx/api/controllers"
	"golang-fx/infrastructure"
)

//UserRoutes -> setting routes for user entity
func UserRoutes(routes infrastructure.GinRouter, logger infrastructure.Logger, controller controllers.UserController) {
	logger.Zap.Info("Setting up user routes")

	userGroup := routes.Gin.Group("/clients-assigned-to")
	{
		userGroup.GET("/", controller.GetAllUser)
	}
}
