package services

import "net/http"

type JobService interface {
	GetJobsHandler(w http.ResponseWriter, req *http.Request)
	GetJobDetailHandler(w http.ResponseWriter, req *http.Request)
}
