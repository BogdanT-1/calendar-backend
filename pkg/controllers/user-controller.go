package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BogdanT-1/calendar-backend/pkg/models"
	"github.com/BogdanT-1/calendar-backend/pkg/token"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	hashedPassword, err := utils.HashPassword(CreateUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	existingUser, _ := models.GetUserByEmail(CreateUser.Email)
	if existingUser.Email != "" {
		w.WriteHeader(http.StatusConflict)
		return
	}
	CreateUser.Password = hashedPassword
	CreateUser.CreateUser()
	res, _ := json.Marshal(CreateUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &models.LoginUserRequest{}
	utils.ParseBody(r, LoginUser)
	user, _ := models.GetUserByEmail(LoginUser.Email)
	if user.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := utils.CheckPassword(LoginUser.Password, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	server, _ := token.ReleaseToken()

	accessToken, err := server.TokenMaker.CreateToken(LoginUser.Email, server.Config.AccessTokenDuration)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := models.LoginUserResponse{
		AccessToken: accessToken,
		Username:    user.Username,
		Email:       user.Email,
	}
	res, _ := json.Marshal(rsp)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
