package routes

import (
	"golang-fx/api/controllers"
	"golang-fx/infrastructure"
)

//UserRoutes -> setting routes for user entity
func UserRoutes(logger infrastructure.Logger, controller controllers.UserController) {
	logger.Zap.Info("Setting up user routes")
	{

		// client.GET("/", controller.)

	}
	clientAssigned := s.Handler.Gin.Group("/clients-assigned-to").Use(s.AuthMiddleware.Handle())
	{
		clientAssigned.GET("/:id", s.Controller.GetAllClientAssigned())
	}
}
