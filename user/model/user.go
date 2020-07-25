package model

type User struct {
	ID       int
	Email    string
	Password string
	Urls     []URL
}