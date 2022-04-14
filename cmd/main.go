package main

import (
	"log"
	"net/http"

	"github.com/BogdanT-1/calendar-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterCalendarRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}