package services

import (
	"github.com/jinzhu/copier"
	"github.com/obed/demo-project/models"
	"github.com/obed/demo-project/repository"
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

func (service *stdService) CreateUser(reqData models.Request) (res models.Response, err error) {
	var userData models.User
	err = copier.Copy(&userData, &reqData)
	if err != nil {
		service.stdRepo.CreateUser(userData)
		if err != nil {
			service.logger.Error(err.Error())
		}
	}
	return res, err
}
