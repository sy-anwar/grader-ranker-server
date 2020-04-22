package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"github.com/sy-anwar/grader-ranker-server/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEventRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}