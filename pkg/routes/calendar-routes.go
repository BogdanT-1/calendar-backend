package routes

import (
	"fmt"

	"github.com/BogdanT-1/calendar-backend/pkg/controllers"
	"github.com/BogdanT-1/calendar-backend/pkg/token"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	config     utils.Config
	tokenMaker token.Maker
	router     *mux.Router
}

var RegisterCalendarRoutes = func(config utils.Config, router *mux.Router) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	router.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/complete", controllers.CompleteTasks).Methods("POST")
	router.HandleFunc("/task/{taskId}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/createUser/", controllers.CreateUser).Methods("POST")

	server.router = router
	return server, nil
}
