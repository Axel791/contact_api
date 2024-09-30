package user

import "github.com/Axel791/contact_api/internal/models"

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int64) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}
