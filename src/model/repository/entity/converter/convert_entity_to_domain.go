package converter

import (
	"sosservice/src/model"
	"sosservice/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Password)

	domain.SetID(entity.ID.Hex())

	return domain
}
