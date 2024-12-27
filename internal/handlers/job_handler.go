package handlers

import (
	"Aurelia/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type JobHandler struct {
	useCase *usecase.JobUsecase
}

func NewJobHandler(useCase *usecase.JobUsecase) *JobHandler {
	return &JobHandler{useCase: useCase}
}

// GET /api/jobs
func (jH *JobHandler) GetJobsHandler(w http.ResponseWriter, req *http.Request) {
	jobs, err := jH.useCase.GetJobs()
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

// GET api/jobs group by job type

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

	jobs, err := jH.useCase.GetJobByID(jobID)
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
