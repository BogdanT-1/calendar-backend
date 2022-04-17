package main

import (
	"log"
	"net/http"

	"github.com/BogdanT-1/calendar-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CORSRouterDecorator struct {
    R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")
    }
    if req.Method == "OPTIONS" {
        return
    }

    c.R.ServeHTTP(rw, req)
}

func main(){
	r := mux.NewRouter()
	routes.RegisterCalendarRoutes(r)
	http.Handle("/", &CORSRouterDecorator{r})
	log.Fatal(http.ListenAndServe("localhost:9010", &CORSRouterDecorator{r}))
}