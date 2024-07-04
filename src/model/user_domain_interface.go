package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	HashPassword()

	SetID(string)
	GetID() string
}

func NewUserDomain(email string, password string, name string, age int) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(email string, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(name string, age int) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
