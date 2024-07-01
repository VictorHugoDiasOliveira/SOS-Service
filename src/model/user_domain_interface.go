package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	HashPassword()

	SetID(string)
	GetID() string
}

func NewUserDomain(email string, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
