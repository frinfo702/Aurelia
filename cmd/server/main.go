package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"Aurelia/internal/handlers"
	"Aurelia/internal/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// TODO: db setup
var (
	dbHost     = "db"
	dbPort     = 5432
	dbUser     = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = "aurelia_db"
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

	logger := logrus.New()

	logger.Debug("This is a debug message")
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	// err = createTable(db) // for only testing in browse!! delete it after!!
	// if err != nil {
	// 	log.Fatal("failed to create table", err)
	// }
	// err = loadInitialData(db) // for only testing in browse!! delete it after!!
	// if err != nil {
	// 	log.Fatal("failed to load initial data", err)
	// }
	r := NewRouter(db)
	log.Println("server started at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTable(db *sql.DB) error {
	query, err := os.ReadFile("migrations/001_init_schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read data %v", err)
	}
	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("failed to excute create query %v", err)
	}

	log.Println("Successfully initialize table")
	return nil
}

func loadInitialData(db *sql.DB) error {
	data, err := os.ReadFile("migrations/realistic_data.sql")
	if err != nil {
		return fmt.Errorf("failed to read data %v", err)
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to excute data %v", err)
	}

	log.Println("Successfully loaded data")
	return nil
}

func NewRouter(db *sql.DB) *mux.Router {

	jobHandler := handlers.NewJobHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("/api/jobs", jobHandler.GetJobsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/jobs/detail", jobHandler.GetJobDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/jobs", jobsHandler)
	r.HandleFunc("/jobs/detail", jobsDetailHandler)
	r.Use(middleware.LoggingMiddleware)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("frontend_mock/static"))),
	)

	return r
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/index.html")
}

func jobsHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/jobs.html")
}

func jobsDetailHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/job_detail.html")
}
