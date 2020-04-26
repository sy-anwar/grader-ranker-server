package routes

import (
	"github.com/gorilla/mux"
	"github.com/sy-anwar/grader-ranker-server/pkg/controllers"
)

var RegisterRankRoutes = func(router *mux.Router) {
	router.HandleFunc("/ranker/{idExam}", controllers.GetRankByIDExam).Methods("GET")
}