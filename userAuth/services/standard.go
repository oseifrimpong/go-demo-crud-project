package service

import (
	"github.com/obed/demo-project/userAuth/model"
	"github.com/obed/demo-project/userAuth/repository"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type stdService struct {
	stdRepo repository.Repository
	logger  *zap.Logger
}

func NewStdService(repo repository.Repository, logger *zap.Logger) Service {
	return &stdService{
		stdRepo: repo,
		logger:  logger,
	}
}

func (service *stdService) CreateUser(reqData model.Request) (res model.Response, err error) {

	var userData model.User
	err = copier.Copy(&userData, &reqData)
	if err != nil {
		service.logger.Error(err.Error())
	}

	res, err = service.stdRepo.CreateUser(userData)
	if err != nil {
		service.logger.Error(err.Error())
	}

	return res, err
}
