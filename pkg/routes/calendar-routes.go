package routes

import (
	"github.com/BogdanT-1/calendar-backend/pkg/controllers"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
	"github.com/gorilla/mux"
)

var RegisterCalendarRoutes = func(config utils.Config, router *mux.Router) {

	router.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/complete", controllers.CompleteTasks).Methods("POST")
	router.HandleFunc("/task/{taskId}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/createUser/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/loginUser/", controllers.LoginUser).Methods("POST")
}
