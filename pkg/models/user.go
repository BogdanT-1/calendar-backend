package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:""json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	Username              string    `json:"username"`
	Email                 string    `json:"email"`
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
