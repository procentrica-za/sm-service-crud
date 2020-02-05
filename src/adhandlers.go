package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handlepostadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Post ad has Been Called!")
		advertisement := PostAdvertisement{}
		err := json.NewDecoder(r.Body).Decode(&advertisement)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to add advertisement ")
			return
		}
		var advertisementposted bool
		var id, message string

		querystring := "SELECT * FROM public.addadvertisement('" + advertisement.UserID + "','" + advertisement.AdvertisementType + "','" + advertisement.EntityID + "','" + advertisement.Price + "','" + advertisement.Description + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&advertisementposted, &id, &message)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to add advertisement")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}

		postAdvertisementResult := PostAdvertisementResult{}
		postAdvertisementResult.AdvertisementPosted = advertisementposted
		postAdvertisementResult.ID = id
		postAdvertisementResult.Message = message

		js, jserr := json.Marshal(postAdvertisementResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to register user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetadvertisement() http.HandlerFunc {
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
			fmt.Fprintf(w, "Unable to process DB Function to get user")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to get user")
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
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleupdateadvertisement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Update User Has Been Called...")
		user := updateUser{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update user")
			return
		}

		var userUpdated bool
		var msg string
		querystring := "SELECT * FROM public.updateuser('" + user.UserID + "','" + user.Username + "','" + user.Password + "','" + user.Name + "','" + user.Surname + "','" + user.Email + "')"
		err = s.dbAccess.QueryRow(querystring).Scan(&userUpdated, &msg)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to update user")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to update user")
			return
		}

		updateUserResult := UpdateUserResult{}
		updateUserResult.UserUpdated = userUpdated
		updateUserResult.Message = msg

		js, jserr := json.Marshal(updateUserResult)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to update user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleremoveadvertisement() http.HandlerFunc {
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
			fmt.Fprintf(w, "Unable to process DB Function to delete user")
			fmt.Println(err.Error())
			fmt.Println("Error in communicating with database to delete user")
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
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}
