package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	server := Server{
		router: mux.NewRouter(),
	}

	// Setup Routes for the server
	server.routes()
	handler := removeTrailingSlash(server.router)

	fmt.Printf("starting server on port 8080 .... \n")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
