package controllers

import (
	"golang-fx/api/repository"
	"golang-fx/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController represents userController entity
type UserController struct {
	repository repository.UserRepository
	logger     infrastructure.Logger
}

//NewUserController returns new user controller
func NewUserController(repository repository.UserRepository, logger infrastructure.Logger) UserController {
	return UserController{
		repository: repository,
		logger:     logger,
	}
}

//GetAllUser -> get all users
func (u UserController) GetAllUser(c *gin.Context) {
	users, _ := u.repository.GetAll()
	c.JSON(http.StatusOK, users)
}
