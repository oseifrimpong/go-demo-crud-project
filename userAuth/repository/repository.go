package repository

import "github.com/obed/demo-project/models"

type Repository interface {
	CreateUser(data models.User) (response models.Response, err error)
	// signIn(data models.User)
}
