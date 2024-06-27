package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	IsAdmin  bool
	// UserInfoID uint
	// UserInfo   UserInfo
}

// type UserInfo struct {
// 	gorm.Model
// 	Name     string
// 	Pix      string `gorm:"unique"`
// 	Cpf      string `gorm:"unique"`
// 	Approved bool
// 	CauseID  uint
// 	Cause    Cause
// }

type Cause struct {
	gorm.Model
	Name string
}
