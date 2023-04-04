package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID      uint    `gorm:"primaryKey"`
	Address string  `gorm:"type:varchar(50);unique"`
	Avatar  *string `gorm:"type:varchar(255);"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

func (m User) Create(user *User) error {
	return DB.Create(&user).Error
}

func (m User) FirstOrCreate(user *User) error {
	return DB.FirstOrCreate(user, User{Address: user.Address}).Error
}

func (m User) GetByID(ID uint) (*User, error) {
	var user User
	err := DB.First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
