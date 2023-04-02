package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    UserName string
	Password string
	Status string
}

const (
	PassWordCost = 12
	Active string = "active"
	Inactive string = "inactive"
)

func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}