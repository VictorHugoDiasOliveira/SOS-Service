package service

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"

	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUserService model",
		zap.String("journey", "DeleteUserService"),
	)

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}
