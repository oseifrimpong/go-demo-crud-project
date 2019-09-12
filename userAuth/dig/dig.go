package dig

import (
	"github.com/obed/demo-project/userAuth/config"
	"github.com/obed/demo-project/userAuth/repository"
	"github.com/obed/demo-project/userAuth/service"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildProject() *dig.Container {
	// container.Provide()
	container.Provide(config.NewDb)
	container.Provide(config.NewConfig)
	container.Provide(repository.NewUserRepo)
	container.Provide(service.NewUserService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
