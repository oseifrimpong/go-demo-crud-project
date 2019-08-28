package delivery

import (
	"github.com/obed/demo-project/userAuth/model"
)

type Delivery interface {
	CreateUser(data model.Request) (response model.Response)
}
