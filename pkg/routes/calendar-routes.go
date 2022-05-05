package routes

import (
	"net/http"
	"strings"

	"github.com/BogdanT-1/calendar-backend/pkg/controllers"
	"github.com/BogdanT-1/calendar-backend/pkg/models"
	"github.com/BogdanT-1/calendar-backend/pkg/token"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
	"github.com/gorilla/mux"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if len(authorizationHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		accessToken := fields[1]
		var tokenMaker token.PasetoMaker
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := r.Header.Get("User")
		models.LoggedInUser = user

		w.Header().Set(authorizationPayloadKey, payload.Username)
		h.ServeHTTP(w, r)
	})
}

var RegisterCalendarRoutes = func(config utils.Config, router *mux.Router) {

	authRouter := router.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/createUser/", controllers.CreateUser).Methods("POST")
	authRouter.HandleFunc("/loginUser/", controllers.LoginUser).Methods("POST")
	authRouter.HandleFunc("/refreshToken/", controllers.RenewToken).Methods("POST")

	taskRouter := router.PathPrefix("/v1").Subrouter()
	taskRouter.Use(Middleware)
	taskRouter.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/day/{assignedDate}", controllers.GetTasksByAssignedDate).Methods("GET")
	taskRouter.HandleFunc("/tasks/complete", controllers.CompleteTasks).Methods("POST")
	taskRouter.HandleFunc("/task/{taskId}", controllers.GetTaskById).Methods("GET")
	taskRouter.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
