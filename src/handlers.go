package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

		w.WriteHeader(200)
		fmt.Fprint(w, "Received Your Json Payload!")

	}
}

func (s *Server) handlerespond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Println("Hit!")
		return
	}
}
