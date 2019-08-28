package service

import (
	"github.com/obed/demo-project/userAuth/model"
)

type Service interface {
	CreateUser(data model.Request) (res model.Response, err error)
}
