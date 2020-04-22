package controllers

import (
	"net/http"
	"encoding/json"
	// "github.com/sy-anwar/grader-ranker-server/pkg/models"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}