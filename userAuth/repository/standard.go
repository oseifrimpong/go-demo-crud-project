package repository

import (
	"encoding/json"

	"github.com/imroc/req"
	"github.com/obed/demo-project/models"
	"go.uber.org/zap"
)

type stdRepository struct {
	logger        *zap.Logger
	serverAddress string
}

func NewStdRepository(logger *zap.Logger) Repository {
	serverAddress := "https//:localhost:8080"

	return &stdRepository{
		logger:        logger,
		serverAddress: serverAddress,
	}
}

func (repo *stdRepository) CreateUser(data models.User) (res models.Response, err error) {
	var newUser models.User
	url := repo.serverAddress + "/createUser"
	header := req.Header{
		"Accept": "application/json",
	}
	r, err := req.Post(url, header, req.BodyJSON(&newUser))
	if err != nil {
		repo.logger.Error(err.Error())
	}
	response := r.Response()
	defer response.Body.Close()

	result := models.Response{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		repo.logger.Error(err.Error())
	}

	res = models.Response{
		Code:   result.Code,
		Result: result.Result,
		Data:   result.Data,
	}
	return
}
