package service

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return err
	}

	return nil
}
