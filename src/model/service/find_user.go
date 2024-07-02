package service

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Starting FindUserByIdServices",
		zap.String("journey", "FindUserByIdService"),
	)

	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Starting FindUserByEmailServices",
		zap.String("journey", "FindUserByEmailService"),
	)
	return ud.userRepository.FindUserByEmail(email)
}
