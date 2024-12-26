package handlers

import "Aurelia/internal/models"

type JobRepository interface {
	SelectJobList(hiringType int, technologyType string, incomeRange int, jobTag string, usedTechnology string) ([]models.Job, error)
	GetJobDetail(jobID int) (*models.Job, error)
}
