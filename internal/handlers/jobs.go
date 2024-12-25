package handlers

import (
	"Aurelia/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

// GET /api/jobs/{id}
func (jH *JobHandler) GetJobDetailHandler(w http.ResponseWriter, req *http.Request) {
	// parse {id} from url request body

	vars := mux.Vars(req)
	jobID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, "Invalid job id", http.StatusNotFound)
	}

	// get the job which corresponds with {id}
	var job models.Job
	job.JobID = jobID
	const query = ` SELECT 
		company_id,
		hiring_type,
		technology_type,
		income_range,
		job_tag,
		requirements,
		used_technology
		FROM jobs	
		WHERE job_id = $1
	`
	row := jH.db.QueryRow(query, jobID)
	if err := row.Err(); err != nil {
		log.Println("failed to fetch job", err)
		http.Error(w, "failed to fetch job", http.StatusNotFound)
		return
	}

	err = row.Scan(
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
			log.Printf("job not found: %v", err)
			http.Error(w, "job not found", http.StatusNotFound)
			return

		} else {
			log.Printf("error scanning job row: %v", err)
			http.Error(w, "failed to scan job data", http.StatusInternalServerError)
			return
		}
	}

	// return job as json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(job); err != nil {
		log.Printf("error encoding response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}
