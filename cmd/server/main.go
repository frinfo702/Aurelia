package main

import (
	"Aurelia/internal/handlers"
	"Aurelia/internal/usecase"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"Aurelia/internal/infrastructure/db"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer dbConn.Close()
	jobRepo := db.NewPostgresJobRepository(dbConn)
	jobUseCase := usecase.NewJobUsecase(jobRepo)
	jobHandler := handlers.NewJobHandler(jobUseCase)

	r := mux.NewRouter()

	r.HandleFunc("/api/jobs", jobHandler.GetJobsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/jobs/{id:[0-9]+}", jobHandler.GetJobByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/jobs", jobs).Methods(http.MethodGet)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(http.ListenAndServe(":8080", r))
	}
}

func jobs(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Here is the place for jobs")
}
