package models

type User struct {
	IDUser   uint32 `json:"id_user"`
	Username string `json:"username"`
	Name     string `json:"name"`
	School   string `json:"school"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type FetchUser struct {
	Data 	UserData  `json:"data"`
}

type UserData struct {
	User	User
}