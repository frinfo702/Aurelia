package usecase

import "Aurelia/internal/domain/models"

// JobUseCase defines job management business logic contract
type JobUsecase interface {
	GetJobs() ([]models.Job, error)
	GetJobByID(id int) (*models.Job, error)
	CreateJob(job *models.Job) error
	DeleteJob(id int) error
}
