package main

import (
	"Aurelia/internal/handlers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// TODO: db setup
var (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = "Aurelia_db"
	psqlInfo   = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
)

func main() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()

	r := NewRouter(db)
	log.Println("server started at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func NewRouter(db *sql.DB) *mux.Router {

	jobHandler := handlers.NewJobHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("api/jobs", jobHandler.GetJobsHandler).Methods(http.MethodGet)
	r.HandleFunc("api/jobs/{id:[0-9]+}", jobHandler.GetJobDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/jobs", jobs).Methods(http.MethodGet)
	log.Println("server started at port: 8080")

	return r
}

func jobs(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Here is the place for jobs")
}
