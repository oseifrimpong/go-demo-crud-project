package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/obed/demo-project/userAuth/config"
	"github.com/obed/demo-project/userAuth/model"
	service "github.com/obed/demo-project/userAuth/services"
	"go.uber.org/zap"
)

type HttpDelivery struct {
	service service.Service
	logger  *zap.Logger
}

func NewGinDelivery(r *gin.Engine, service service.Service, logger *zap.Logger) {
	delivery := &HttpDelivery{
		service,
		logger,
	}

	v1 := r.Group("/api/v1")
	{
		// v1.GET("/test", delivery.Test)
		v1.POST("createuser", delivery.SignUp)
	}
}

func (a *HttpDelivery) SignUp(c *gin.Context) {
	var userData model.User
	var request model.Request

	c.ShouldBindJSON(&userData)
	// a.logger.Error(err.Error())
	if config.Database.Create(&userData).Error != nil {
		c.SecureJSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
	} else {
		response := a.CreateUser(request)
		c.SecureJSON(response.Code, response)
		c.SecureJSON(http.StatusOK, gin.H{"Status": "User registered Successfully"})
	}

	// strRequest, _ := json.Marshal(&request)
	// a.logger.Debug(string(strRequest))
}

func (a *HttpDelivery) CreateUser(ctx model.Request) (response model.Response) {
	response, _ = a.service.CreateUser(ctx)

	return response
}
