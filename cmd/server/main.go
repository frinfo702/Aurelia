package main

import (
	"Aurelia/cmd/server/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = 5432
	dbUser     = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	psqlInfo   = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName)
)

func main() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
		return
	}
	defer db.Close()

	r := router.NewRouter(db)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(http.ListenAndServe(":8080", r))
	}

}
