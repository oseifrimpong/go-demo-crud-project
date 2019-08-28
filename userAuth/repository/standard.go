package repository

import (
	"encoding/json"

	"github.com/imroc/req"
	"github.com/obed/demo-project/userAuth/model"
	"go.uber.org/zap"
)

type stdRepository struct {
	logger        *zap.Logger
	serverAddress string
}

func NewStdRepository(logger *zap.Logger) Repository {
	serverAddress := "https://localhost:8080"

	return &stdRepository{
		logger:        logger,
		serverAddress: serverAddress,
	}
}

func (repo *stdRepository) CreateUser(data model.User) (res model.Response, err error) {
	url := repo.serverAddress + "/createuser"
	header := req.Header{
		"Accept": "applications/json",
	}
	r, err := req.Post(url, header)
	if err != nil {
		repo.logger.Error(err.Error())
	}

	resp := r.Response()
	defer resp.Body.Close()

	result := model.Response{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		repo.logger.Error(err.Error())
	}

	res = model.Response{
		Code:    result.Code,
		Data:    result.Data,
		Message: result.Message,
	}
	return
}
