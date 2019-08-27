package delivery

import "github.com/obed/demo-project/models"

type Delivery interface {
	CreateUser(userData models.Request) (response models.Response)
}
