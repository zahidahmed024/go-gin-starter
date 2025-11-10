package repositories

import (
	"go-gin-starter/internal/models"
	"go-gin-starter/pkg/database"
)

// UserRepository handles the database operations for the User model
type UserRepository struct{}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetUserByEmail retrieves a user by email from the database
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
