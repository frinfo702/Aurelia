package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/usecase"

	"github.com/gorilla/mux"
)

type JobHandler struct {
	jobUC usecase.JobUsecase
}

func NewJobHandler(jobUC usecase.JobUsecase) *JobHandler {
	return &JobHandler{jobUC: jobUC}
}

// GET /api/jobs
func (jH *JobHandler) GetJobsHandler(w http.ResponseWriter, req *http.Request) {
	jobs, err := jH.jobUC.GetJobs()
	if err != nil {
		log.Printf("failed to fetch job list: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to fetch job list"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Printf("error encoding response: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to encode response"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}
}

// GET /api/jobs/{id}
func (jH *JobHandler) GetJobByIDHandler(w http.ResponseWriter, req *http.Request) {
	// parse {id} from url request body

	jobID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		log.Printf("failed to parse job id: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "invalid job	id"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
	}

	jobs, err := jH.jobUC.GetJobByID(jobID)
	if err != nil {
		log.Printf("failed to fetch job list: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to fetch job list"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Printf("error encoding response: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to encode response"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

}

// POST /api/jobs
func (jH *JobHandler) CreateJobHandler(w http.ResponseWriter, req *http.Request) {
	// parse request json body to job struct
	var job *models.Job
	err := json.NewDecoder(req.Body).Decode(&job)
	if err != nil {
		log.Printf("failed to parse job: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "invalid job"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

	// create job
	err = jH.jobUC.CreateJob(job)
	if err != nil {
		log.Printf("failed to create job: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to create job"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

	insertedJobID, err := jH.jobUC.GetJobByID(job.JobID)
	if err != nil {
		log.Printf("failed to fetch job: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "failed to fetch job"})
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

	log.Printf("job created: %v", insertedJobID)
	log.Println("Notification sent") // placeholder for notification
	w.WriteHeader(http.StatusCreated)
}
