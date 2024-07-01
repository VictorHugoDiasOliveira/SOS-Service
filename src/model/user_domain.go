package model

type userDomain struct {
	id       string
	email    string
	password string
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
