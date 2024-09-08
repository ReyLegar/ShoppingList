package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
