package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/sy-anwar/grader-ranker-server/pkg/config"
)


type Session struct {
	EventID			uint32	`gorm:"primary_key;auto_increment:false" json:"id_event"`
	IDSession 		uint32	`gorm:"primary_key;auto_increment:false" json:"id_session"`
	SessionScore	int		`json:"session_score"`
}


func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Session{})
}

func (session *Session) CreateSession() *Session {
	// Update if record found else Insert
	db.Where(Session{
		EventID: session.EventID,
		IDSession: session.IDSession,
	}).Assign(Session{
		SessionScore: session.SessionScore,
	}).FirstOrCreate(&Session{})
	db.Where(Session{
		EventID: session.EventID,
		IDSession: session.IDSession,
	}).First(&session)
	return session
}


