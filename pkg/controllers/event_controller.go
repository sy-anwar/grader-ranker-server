package controllers

import (
	"log"
	"errors"
	"net/http"
	"encoding/json"
	"github.com/sy-anwar/grader-ranker-server/pkg/models"
	"github.com/sy-anwar/grader-ranker-server/pkg/utils"
)

func CreateEvents(w http.ResponseWriter, r *http.Request) {
	var data models.Examsheets
	err := utils.DecodeJSONBody(w, r, &data)
    if err != nil {
        var mr *utils.MalformedRequest
        if errors.As(err, &mr) {
            http.Error(w, mr.Msg, mr.Status)
        } else {
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
	}
	if utils.CheckNilFieldsExamsheets(data) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp := models.CreateEvents(data)
	res, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Hello")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}