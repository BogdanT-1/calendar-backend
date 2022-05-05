package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BogdanT-1/calendar-backend/pkg/models"
	"github.com/BogdanT-1/calendar-backend/pkg/token"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func RenewToken(w http.ResponseWriter, r *http.Request) {
	req := &renewAccessTokenRequest{}
	utils.ParseBody(r, req)

	server, _ := token.ReleaseToken()
	refreshPayload, err := server.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session, _ := models.GetSessionByToken(req.RefreshToken)
	if session.IsBlocked {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if session.Username != refreshPayload.Username {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, err := server.TokenMaker.CreateToken(session.Username, server.Config.AccessTokenDuration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Now().Add(time.Minute * 15),
	}

	res, _ := json.Marshal(rsp)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
