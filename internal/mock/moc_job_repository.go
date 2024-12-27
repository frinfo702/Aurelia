package mock

import (
	"Aurelia/internal/domain/models"

	"github.com/stretchr/testify/mock"
)

// Create a mock repository
type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) GetJobDetail(jobID int) (*models.Job, error) {
	args := m.Called(jobID)
	return args.Get(0).(*models.Job), args.Error(1)
}
