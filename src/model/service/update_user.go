package service

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUserService model",
		zap.String("journey", "UpdateUserService"),
	)

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return err
	}

	return nil
}
