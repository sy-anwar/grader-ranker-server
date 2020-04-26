package models

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
)

type UserRank struct {
	User
	IDExam	uint32	`json:"id_exam"`
	Rank	uint32	`json:"rank"`
}

type Rank struct {
	Data []UserRank `json:"data"`
}

func ConcateUserRank(events []Event) []UserRank {
	var rank []UserRank
	for i, _ := range events {
		idUser := events[i].IDUser
		urlGet := fmt.Sprintf("https://dummy.edukasystem.id/user/%d", idUser)
		client := &http.Client{ Timeout: time.Second * 10}
		res, err := client.Get(urlGet)
		if err != nil {
			log.Println("Get User, ID %d failed", idUser)
		} 
		var fetchUser FetchUser
		json.NewDecoder(res.Body).Decode(&fetchUser)
		userRank := UserRank {
			User   : fetchUser.Data.User,
			IDExam : events[i].IDExam,
			Rank   : uint32(i) + 1,
		}
		rank = append(rank, userRank)
	}

	return rank
}