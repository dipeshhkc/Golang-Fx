package repository

import (
	"golang-fx/api/models"
	"golang-fx/infrastructure"
)

//UserRepository -> repository to user entity
type UserRepository struct {
	db infrastructure.Database
}

//NewUserRepository -> returns new instance of user repository
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

//GetAll -> Get All User
func (u UserRepository) GetAll() (users []models.User, err error) {
	return users, u.db.DB.Find(&users).Error
}
