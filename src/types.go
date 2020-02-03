package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	dbAccess *sql.DB
	router   *mux.Router
}

type UserID struct {
	UserID string `json:userid`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type getUser struct {
	UserID   string `json:id`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type updateUser struct {
	UserID   string `json:id`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}

type UpdateUserResult struct {
	UserUpdated bool `json:"userupdated"`
}

type DeleteUserResult struct {
	UserDeleted bool `json:"userdeleted"`
}

type RegisterUserResult struct {
	UserCreated bool   `json:"usercreated"`
	Username    string `json:"username"`
	UserID      string `json:"userid"`
}

type Config struct {
	UserName     string
	Password     string
	DatabaseName string
	Port         string
	PostgresHost string
	PostgresPort string
}
