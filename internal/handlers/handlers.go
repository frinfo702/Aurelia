package handlers

import (
	"database/sql"
	"log"
)

var globalDB *sql.DB

// SetDB sets the global database connection for handlers
func SetDB(db *sql.DB) {
	globalDB = db
	log.Println("global db connection set", globalDB)
}

// 他のハンドラも同様に設定
