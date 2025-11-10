package services

import (
	"errors"
	"time"

	"go-gin-starter/internal/models"
	"go-gin-starter/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"go-gin-starter/config"
)

// AuthService handles the business logic for authentication
type AuthService struct {
	userRepository *repositories.UserRepository
}

// NewAuthService creates a new AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		userRepository: &repositories.UserRepository{},
	}
}

// RegisterUser handles the user registration logic
func (s *AuthService) RegisterUser(name, email, password string) (*models.User, error) {
	// Check if user already exists
	_, err := s.userRepository.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser handles the user login logic
func (s *AuthService) LoginUser(email, password string) (string, error) {
	// Get user by email
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	cfg, _ := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
