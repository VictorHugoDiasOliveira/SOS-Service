package service

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating User Service", zap.String("journey", "CreateUserService"))
	if user, _ := ud.FindUserByEmailService(userDomain.GetEmail()); user != nil {
		err := rest_err.NewBadRequestError("This email is already used")
		return nil, err
	}

	userDomain.HashPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	logger.Info("User Service Created Successfully", zap.String("journey", "CreateUserService"))
	return userDomainRepository, nil
}
