package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/obed/demo-project/models"
	"github.com/obed/demo-project/services"
	"go.uber.org/zap"
)

type HttpDelivery struct {
	service services.Service
	logger  *zap.Logger
}

func NewGinDelivery(r *gin.Engine, service services.Service, logger *zap.Logger) {
	delivery := &HttpDelivery{
		service,
		logger,
	}

	v1 := r.Group("api/v1/")
	{
		v1.POST("/createUser", delivery.createUser)
	}
}

func (a *HttpDelivery) CreateUser(context models.Request) (response models.Response) {
	response, _ = a.service.CreateUser(context)
	return response
}
