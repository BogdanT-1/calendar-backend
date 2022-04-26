package controllers

import (
	"net/http"

	"github.com/BogdanT-1/calendar-backend/pkg/models"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	CreateUser.CreateUser()
	w.WriteHeader(http.StatusCreated)
}
