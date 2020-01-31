package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

type User struct {
	Username string `json:"username"`
	Password string `json:"string"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
