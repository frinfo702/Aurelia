package usecase

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/repository/postgresql"
)

type JobUsecase struct {
	Repo postgresql.JobRepository
}

func NewJobUsecase(repo postgresql.JobRepository) *JobUsecase {
	return &JobUsecase{Repo: repo}
}

func (uc *JobUsecase) GetJobs() ([]models.Job, error) {
	return uc.Repo.FindAll()
}

func (uc *JobUsecase) GetJobByID(id int) (*models.Job, error) {
	return uc.Repo.FindByID(id)
}

// TODO: Implement the following methods
/*
func (uc *JobUseCase) CreateJob(job *models.Job) error {
    return uc.repo.Save(job)
}

func (uc *JobUseCase) DeleteJob(id int) error {
    return uc.repo.Delete(id)
}
*/
