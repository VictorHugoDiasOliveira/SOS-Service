package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	HashPassword()
}

func NewUserDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email, password,
	}
}

type userDomain struct {
	email    string
	password string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) HashPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
