package model

type User struct {
	Id int `json:"id,ok"`
	Username string `json:"username,ok"`
	Password string `json:"password,ok"`
}