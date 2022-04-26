package token

import (
	"fmt"
	"log"

	"github.com/BogdanT-1/calendar-backend/pkg/utils"
)

type Server struct {
	Config     utils.Config
	TokenMaker Maker
}

func ReleaseToken() (*Server, error) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	tokenMaker, err := NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	ServerConf := &Server{
		Config:     config,
		TokenMaker: tokenMaker,
	}

	return ServerConf, nil
}
