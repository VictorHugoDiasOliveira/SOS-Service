package service

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
)

func (ud *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {

	userDomain.HashPassword()

	user, err := ud.FindUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
