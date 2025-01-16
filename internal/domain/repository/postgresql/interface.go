package postgresql

import "Aurelia/internal/domain/models"

type JobRepository interface {
	FindAll() ([]models.Job, error)
	FindByID(id int) (*models.Job, error)
	Insert(job *models.Job) error
	Delete(id int) error
}

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByEmail(email string) (*models.User, error)
	Insert(user *models.User) error
	CheckPassword(hash, password string) bool
}
