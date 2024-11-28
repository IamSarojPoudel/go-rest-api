package services

import (
	"errors"
	"rest-api/internal/middleware"
	"rest-api/internal/models"
	"rest-api/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}
func (s *authService) Register(user *models.User) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	if err := s.userRepo.Create(user); err != nil {
		return "", errors.New("user already exist")
	}
	token, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
