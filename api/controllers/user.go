package controllers

import (
	"golang-fx/api/repository"
	"golang-fx/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repository repository.UserRepository
	logger     infrastructure.Logger
}

func NewUserController(repository repository.UserRepository, logger infrastructure.Logger) UserController {
	return UserController{
		repository: repository,
		logger:     logger,
	}
}

func (u UserController) getAllUser(c *gin.Context) {
	users, _ := u.repository.GetAll()
	c.JSON(http.StatusOK, users)
}
