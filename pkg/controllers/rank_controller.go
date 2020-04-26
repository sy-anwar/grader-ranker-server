package controllers

import (
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sy-anwar/grader-ranker-server/pkg/models"
)

func GetRankByIDExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idExam64, err := strconv.ParseUint(params["idExam"], 10, 64)
	idExam32 := uint32(idExam64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	events := models.GetEventsByIDExam(idExam32)
	var response models.Rank
	response.Data = models.ConcateUserRank(events)
	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}