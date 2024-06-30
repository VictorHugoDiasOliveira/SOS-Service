package model

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model",
		zap.String("journey", "createUser"),
	)

	ud.HashPassword()

	return nil
}
