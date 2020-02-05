package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

var dbconf dbConfig
var conf Config

func init() {
	dbconf = CreateDbConfig()
	conf = CreateConfig()
}

func CreateDbConfig() dbConfig {
	dbconf := dbConfig{
		UserName:     os.Getenv("POSTGRES_USERNAME"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName: os.Getenv("DATABASENAME"),
		Port:         os.Getenv("PORT"),
		PostgresHost: os.Getenv("POSTGRESHOST"),
		PostgresPort: os.Getenv("POSTGRESPORT"),
	}
	return dbconf

}

func CreateConfig() Config {
	conf := Config{
		ListenServePort: os.Getenv("LISTEN_AND_SERVE_PORT"),
	}
	return conf
}

func openDatabase(host string, port string, user string, password string, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connecting to database")
	}

	err = db.Ping()
	for retry := 0; err != nil && retry < 20; err = db.Ping() {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error: ", err.Code.Name())
		}
		fmt.Println("Sleeping till connection opens")
		time.Sleep(1 * time.Second)
		retry++
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to DB has been opened.")

	return db, err
}

func main() {
	conn, err := openDatabase(dbconf.PostgresHost, dbconf.PostgresPort, dbconf.UserName, dbconf.Password, dbconf.DatabaseName)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Could not establish connection to the database...")
	}

	server := Server{
		dbAccess: conn,
		router:   mux.NewRouter(),
	}

	// Setup Routes for the server
	server.routes()
	handler := removeTrailingSlash(server.router)

	fmt.Printf("starting server on port " + conf.ListenServePort + " .... \n")
	log.Fatal(http.ListenAndServe(":"+conf.ListenServePort, handler))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
