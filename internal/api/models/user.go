package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Address string  `gorm:"type:varchar(50);unique"`
	Avatar  *string `gorm:"type:varchar(255);"`
}

func (m User) Create(user *User) (*User, error) {
	return user, DB.Create(&user).Error
}

func (m User) FirstOrCreate(user *User) (*User, error) {
	return user, DB.FirstOrCreate(&user).Error
}
