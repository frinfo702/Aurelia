package router

import (
	"Aurelia/cmd/server/router/middleware"
	"Aurelia/internal/domain/repository/postgresql"
	"Aurelia/internal/domain/usecase"
	"Aurelia/internal/handlers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	jobRepo := postgresql.NewJobRepository(db)
	useCase := usecase.NewJobUsecase(jobRepo)
	jobHandler := handlers.NewJobHandler(useCase)

	userRepo := postgresql.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	authHandler := handlers.NewAuthHandler(authUC)

	r := mux.NewRouter()

	r.HandleFunc("/api/jobs", jobHandler.GetJobsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/jobs/{id:[0-9]+}", jobHandler.GetJobByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/jobs", jobHandler.CreateJobHandler).Methods(http.MethodPost)

	// authorization
	r.HandleFunc("/api/auth/signup", authHandler.SignUpHander).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/login", authHandler.LoginHandler).Methods(http.MethodPost)

	sub := r.PathPrefix("api/comapanies/me").Subrouter()
	sub.Use(middleware.ValidateJWTMiddleware)
	// sub.HandleFunc("/jobs", CompanyJobsHandler).Methods(http.MethodGet)

	// frontend
	r.HandleFunc("/", htmlHomeHandler)
	r.HandleFunc("/jobs", htmlJobsHandler)
	r.HandleFunc("/jobs/{id:[0-9]+}", htmlJobsDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/apply", htmlApplyJobHandler).Methods(http.MethodGet)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("frontend_mock/static"))),
	)

	return r
}

// /
func htmlHomeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/index.html")
}

// /jobs
func htmlJobsHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/jobs.html")
}

// /jobs/{id:[0-9]+}
func htmlJobsDetailHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/job_detail.html")
}

// /apply
func htmlApplyJobHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "frontend_mock/apply_job.html")
}
