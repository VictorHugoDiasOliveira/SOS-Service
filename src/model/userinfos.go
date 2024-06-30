package model

import "gorm.io/gorm"

type UserInfos struct {
	Id       uint   `gorm:"primary key;autoIncrement" json:"id"`
	Name     string `gorm:"not null;size:100" json:"name"`
	Pix      string `gorm:"unique;not null" json:"pix"`
	Cpf      string `gorm:"unique;not null" json:"cpf"`
	Approved bool   `gorm:"default:false" json:"approved"`
}

func MigrateUserInfos(db *gorm.DB) error {
	err := db.AutoMigrate(&UserInfos{})
	return err
}
