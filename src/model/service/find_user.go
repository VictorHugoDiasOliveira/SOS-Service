package service

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
)

func (ud *userDomainService) FindUserByIdService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserByEmail(email)
}
