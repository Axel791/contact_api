package user

import (
	"github.com/Axel791/contact_api/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Методы пользователей

func (r *Repository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Contacts").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Preload("Contacts").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
