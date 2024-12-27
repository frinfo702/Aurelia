package usecase

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/repository"
)

type JobUsecase struct {
	repo repository.JobRepository
}

func NewJobUsecase(repo repository.JobRepository) *JobUsecase {
	return &JobUsecase{repo: repo}
}

func (uc *JobUsecase) GetJobs() ([]models.Job, error) {
	return uc.repo.FindAll()
}

func (uc *JobUsecase) GetJobByID(id int) (*models.Job, error) {
	return uc.repo.FindByID(id)
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
