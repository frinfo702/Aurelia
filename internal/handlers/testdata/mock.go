package testdata

import (
	"Aurelia/internal/domain/models"

	"github.com/stretchr/testify/mock"
)

// これをRepositoryとして認識させる(interfaceを使う)
type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) FindAll() ([]models.Job, error) {
	args := m.Called()
	return args.Get(0).([]models.Job), args.Error(1)
}

// FindByID
func (m *MockJobRepository) FindByID(id int) (*models.Job, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Job), args.Error(1)
}

// Save
func (m *MockJobRepository) Insert(job *models.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

// Delete
func (m *MockJobRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
