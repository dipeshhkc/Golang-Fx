package routes

import (
	"golang-fx/api/controllers"
	"golang-fx/infrastructure"
)

// UserRoutes model
type UserRoutes struct {
	Logger     infrastructure.Logger
	Handler    infrastructure.GinRouter
	Controller controllers.UserController
}

// NewUserRoutes creates a new health check routes
func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.GinRouter,
	controller controllers.UserController,
) UserRoutes {
	return UserRoutes{
		Logger:     logger,
		Handler:    handler,
		Controller: controller,
	}
}

// Setup sets up routes
func (s UserRoutes) Setup() {
	s.Logger.Zap.Info("Setting up user routes")
	client := s.Handler.Gin.Group("/users")
	{
		client.GET("/", s.Controller.GetAllUser)

	}

}
