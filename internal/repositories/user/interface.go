package user

import "github.com/Axel791/contact_api/internal/models"

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}
