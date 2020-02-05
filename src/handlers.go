package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handleloginuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Login User Has Been Called...")
		getusername := r.URL.Query().Get("username")
		getpassword := r.URL.Query().Get("password")
		userLogin := UserLogin{}

		userLogin.Username = getusername
		userLogin.Password = getpassword

		var userid, username string
		querystring := "SELECT * FROM public.loginuser('" + userLogin.Username + "','" + userLogin.Password + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&userid, &username)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}
		loginUserResult := LoginUserResult{}
		if userid == "" {
			loginUserResult.UserLoggedIn = false
			loginUserResult.UserID = ""
			loginUserResult.Message = "Wrong username or password combination!"
		} else {
			loginUserResult.UserLoggedIn = true
			loginUserResult.UserID = userid
			loginUserResult.Message = "Welcome! " + username
		}

		js, jserr := json.Marshal(loginUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handledeleteuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Delete User Has Been Called..")
		getuserid := r.URL.Query().Get("id")
		userid := UserID{}
		userid.UserID = getuserid

		var userDeleted bool
		querystring := "SELECT * FROM public.deleteuser('" + userid.UserID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&userDeleted)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		deleteUserResult := DeleteUserResult{}
		deleteUserResult.UserDeleted = userDeleted
		deleteUserResult.UserID = getuserid

		if userDeleted {
			deleteUserResult.Message = "User Successfully Deleted!"
		} else {
			deleteUserResult.Message = "Unable to Delete User!"
		}

		js, jserr := json.Marshal(deleteUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handleupdateuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Update User Has Been Called...")
		user := updateUser{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}

		var userUpdated bool
		var msg string
		querystring := "SELECT * FROM public.updateuser('" + user.UserID + "','" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userUpdated, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		updateUserResult := UpdateUserResult{}
		updateUserResult.UserUpdated = userUpdated
		updateUserResult.Message = msg

		js, jserr := json.Marshal(updateUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleregisteruser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Register User Has Been Called!")
		user := User{}
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided...")
			return
		}
		var userCreated string
		var username, userid, msg string

		querystring := "SELECT * FROM public.registeruser('" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userCreated, &username, &userid, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}

		regUserResult := RegisterUserResult{}
		regUserResult.UserCreated = userCreated
		regUserResult.Username = username
		regUserResult.UserID = userid
		regUserResult.Message = msg

		js, jserr := json.Marshal(regUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetuser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get User Has Been Called...")
		getuserid := r.URL.Query().Get("id")
		userid := UserID{}
		userid.UserID = getuserid

		var id, username, name, surname, email string

		querystring := "SELECT * FROM public.getuser('" + userid.UserID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &username, &name, &surname, &email)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			fmt.Println(err.Error())
			return
		}
		//fmt.Println("This is User!: " + id)
		user := getUser{}
		user.UserID = id
		user.Username = username
		user.Name = name
		user.Surname = surname
		user.Email = email

		js, jserr := json.Marshal(user)
		fmt.Println(js)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
