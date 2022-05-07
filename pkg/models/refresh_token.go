package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Sessions struct {
	gorm.Model
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (b *Sessions) CreateSession() *Sessions {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetSessionByToken(refreshToken string) (*Sessions, *gorm.DB) {
	var getSession Sessions
	db := db.Where("refresh_token=?", refreshToken).Find(&getSession)
	return &getSession, db
}
