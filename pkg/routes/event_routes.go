package routes

import (
	"github.com/gorilla/mux"
	"github.com/sy-anwar/grader-ranker-server/pkg/controllers"
)

var RegisterEventRoutes = func(router *mux.Router) {
	router.HandleFunc("/grader/event", controllers.GetEvents).Methods("GET")
}