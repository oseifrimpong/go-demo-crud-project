package services

import (
	"github.com/obed/demo-project/models"
)

type Service interface {
	CreateUser(data models.Request) (res models.Response, err error)
}
