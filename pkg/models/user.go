package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:""json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	SessionID             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  User      `json:"user"`
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Email=?", email).Find(&getUser)
	return &getUser, db
}
