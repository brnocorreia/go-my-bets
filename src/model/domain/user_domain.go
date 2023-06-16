package model

import "gorm.io/gorm"

type userDomain struct {
	GeneralInfo gorm.Model
	Email       string
	Name        string
	Nickname    string
	Password    string
}
