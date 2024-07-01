package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	HashPassword()

	GetJSONValue() (string, error)

	SetID(string)
}

func NewUserDomain(email, password string) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
	}
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetJSONValue() (string, error) {
	marshedJson, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(marshedJson), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) HashPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
