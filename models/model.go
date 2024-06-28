package models

//hello
type User struct {
	ID       int    `json:"ID"`
	FULLNAME string `json:"FULLNAME"`
	PHONE    string `json:"PHONE"`
	EMAIL    string `json:"EMAIL"`
	CITY     string `json:"CITY"`
}
