package models

import (
	"reflect"
	"github.com/jinzhu/gorm"
	"github.com/sy-anwar/grader-ranker-server/pkg/config"
)

var db *gorm.DB

type Event struct {
	IDEvent		uint		`gorm:"primary_key" json:"id_event"`
	IDExam      uint      	`json:"id_exam"`
	IDUser      uint       	`json:"id_user"`
	EventScore	int			`json:"event_score"`
	Sessions 	[]Session	`gorm:"foreignkey:EventID" json:"sessions"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Event{})

}

func (event *Event) CreateEvent() *Event {
	// Update if record found else Insert
	db.Where(Event{
		IDExam: event.IDExam,
		IDUser: event.IDUser,
	}).Assign(Event{
		EventScore: event.EventScore,
		Sessions  : event.Sessions,
	}).FirstOrCreate(&Event{})
	db.Where(Event{
		IDExam: event.IDExam,
		IDUser: event.IDUser,
	}).First(&event)
	return event
}

func CreateEvents(examsheet Examsheets) []Event {
	var events []Event
	for i, _ := range examsheet.Data {
		eventScore, sessions := GetEventScoreSessions(examsheet.Data[i].Answersheets)
		event := Event {
			IDExam 		: *(examsheet.Data[i].IDExam),
			IDUser 		: *(examsheet.Data[i].IDUser),
			EventScore  : eventScore,
			Sessions 	: sessions,
		}
		event.CreateEvent()
		events = append(events, event)
	}
	return events
}

func GetEventScoreSessions(answersheets []Answersheet) (int, []Session) {
	eventScore := 0
	var sessions []Session
	for i, _ := range answersheets {
		sessionScore := 0
		answers := reflect.ValueOf(answersheets[i].Answers)
		for i := 0; i < answers.NumField(); i++ {
			val := answers.Field(i).Elem()
			if val.Interface() == true {
				sessionScore++
			}
		}
		session := Session {
			IDSession    : *(answersheets[i].IDSession),
			SessionScore : sessionScore,
		}
		session.CreateSession()
		sessions = append(sessions, session)
		eventScore += sessionScore
	}
	return eventScore, sessions
}
