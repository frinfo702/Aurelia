package usecase

import "Aurelia/internal/domain/models"

type AuthUsecase interface {
	SignUp(user *models.User) error
	Login(email, password string) (string, error)
}
