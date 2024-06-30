package view

import (
	"sosservice/src/controller/model/response"
	"sosservice/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Email: userDomain.GetEmail(),
	}
}
