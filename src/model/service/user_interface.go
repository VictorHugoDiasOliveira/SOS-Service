package service

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
	"sosservice/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIdService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr

	LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
