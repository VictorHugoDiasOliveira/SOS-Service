package repository

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
