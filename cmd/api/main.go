package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/sy-anwar/grader-ranker-server/pkg/config"
	"github.com/sy-anwar/grader-ranker-server/pkg/models"
	"github.com/sy-anwar/grader-ranker-server/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s --> %s %s\n", r.RemoteAddr, r.Method, r.URL)
			lrw := models.NewLoggingResponseWriter(w)

			handler.ServeHTTP(lrw, r)
			statusCode := lrw.StatusCode
        	log.Printf("%s <-- %d %s", r.RemoteAddr, statusCode, http.StatusText(statusCode))
		})
}

func main() {
	// config.ConnectDB()
	r := mux.NewRouter()
	routes.RegisterEventRoutes(r)
	routes.RegisterRankRoutes(r)
	routeLogging := logRequest(r)
	http.Handle("/", routeLogging)
	log.Println("Server started on localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	} 
}
