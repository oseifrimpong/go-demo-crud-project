package repository

import (
	"github.com/obed/demo-project/userAuth/model"
)

type Repository interface {
	CreateUser(data model.User) (res model.Response, err error)
}
