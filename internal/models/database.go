package models

type UserDatabase struct {
	id       string
	username string
	password string
}

type ToDoObject struct {
	Id       string
	TodoText string
	IdUser   string
}
