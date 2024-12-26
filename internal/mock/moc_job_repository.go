package mock

import (
	"Aurelia/internal/models"

	"github.com/stretchr/testify/mock"
)

// mockを作りたい

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) SelectJobList(hiringType int, technologyType string, incomeRange int, jobTag string, usedTechnology string) ([]models.Job, error) {
	args := m.Called(hiringType, technologyType, incomeRange, jobTag, usedTechnology)
	return args.Get(0).([]models.Job), args.Error(1)
}

func (m *MockJobRepository) GetJobDetail(jobID int) (*models.Job, error) {
	args := m.Called(jobID)
	return args.Get(0).(*models.Job), args.Error(1)
}
