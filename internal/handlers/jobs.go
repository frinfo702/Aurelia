package handlers

import (
	"Aurelia/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type JobHandler struct {
	db *sql.DB
}

func NewJobHandler(db *sql.DB) *JobHandler {
	return &JobHandler{db: db}
}

// GET api/jobs?type={} (e.g. fulltime)
func (jH *JobHandler) GetJobsHandler(w http.ResponseWriter, req *http.Request) {
	jobType := req.URL.Query().Get("type")
	const query = `
		SELECT job_id, company_id, hiring_type, technology_type, income_range, 
		job_tag, requirements, used_technology	
		FROM jobs
		WHERE ($1 = '' OR hiring_type = $1)
	`
	rows, err := jH.db.Query(query, jobType)
	if err != nil {
		log.Printf("failed to query jobs: %v", err)
		http.Error(w, "failed to fetch job data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var jobList []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(
			&job.JobID,
			&job.CompanyID,
			&job.HiringType,
			&job.TechnologyType,
			&job.IncomeRange,
			&job.JobTag,
			&job.Requirements,
			&job.UsedTechnology,
		)
		if err != nil {
			log.Printf("error scanning job row: %v", err)
			return
		}
		jobList = append(jobList, job)
	}

	w.Header().Set("Content-Type", "application/json") // look into after. I don't understand
	json.NewEncoder(w).Encode(jobList)
}

// other Handlers for job
func (jH *JobHandler) GetJobDetailHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var jobID int
	if id, ok := queryMap["id"]; ok && len(id) > 0 {
		var err error
		jobID, err = strconv.Atoi(id[0])
		if err != nil {
			log.Printf("Query parameter might be invalid: %v", err)
			http.Error(w, "Invalid job ID provided in query parameter", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Missing job ID", http.StatusBadRequest)
		return
	}

	var job models.Job
	query := `SELECT job_id, company_id, hiring_type, technology_type, income_range, job_tag, requirements, used_technology FROM jobs WHERE job_id = $1`
	err := jH.db.QueryRow(query, jobID).Scan(
		&job.JobID,
		&job.CompanyID,
		&job.HiringType,
		&job.TechnologyType,
		&job.IncomeRange,
		&job.JobTag,
		&job.Requirements,
		&job.UsedTechnology,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Job not found", http.StatusNotFound)
		} else {
			log.Printf("error querying job detail: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
