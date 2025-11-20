package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rafaelchr/task-management/models"
	"github.com/rafaelchr/task-management/repositories"
	"github.com/rafaelchr/task-management/utils"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	// check email
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("Email already registered")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}
