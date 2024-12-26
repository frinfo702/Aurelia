package handlers

import "database/sql"

var globalDB *sql.DB

// SetDB sets the global database connection for handlers
func SetDB(db *sql.DB) {
	globalDB = db
}

// 他のハンドラも同様に設定
