package models

import "time"

type RefreshToken struct {
	ID        uint `gorm:"primaryKey;"`
	UserID    uint
	Token     string `gorm:"type:varchar(512);"`
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m RefreshToken) Create(userID uint, token string, expire time.Time) error {
	err := DB.Create(&RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiredAt: expire,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error

	return err
}

func (m RefreshToken) GetByRefreshToken(token string) (*RefreshToken, error) {
	var refreshToken RefreshToken
	err := DB.First(&refreshToken, RefreshToken{Token: token}).Error

	return &refreshToken, err
}
