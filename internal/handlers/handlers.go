package handlers

import "database/sql"

var globalDB *sql.DB

// SetDB sets the global database connection for handlers
func SetDB(db *sql.DB) {
	globalDB = db
}

type JobHandler struct {
	db *sql.DB
}

func NewJobHandler() *JobHandler {
	return &JobHandler{db: globalDB}
}

// 他のハンドラも同様に設定
