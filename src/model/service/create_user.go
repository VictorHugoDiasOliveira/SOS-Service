package service

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	if user, _ := ud.FindUserByEmailService(userDomain.GetEmail()); user != nil {
		err := rest_err.NewBadRequestError("This email is already used")
		return nil, err
	}

	userDomain.HashPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
