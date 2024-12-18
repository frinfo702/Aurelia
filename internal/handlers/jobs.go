package handlers

import (
	"Aurelia/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type JobHandler struct {
	db *sql.DB
}

func NewJobHandler(db *sql.DB) *JobHandler {
	return &JobHandler{db: db}
}

// GET api/jobs group by job type
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
