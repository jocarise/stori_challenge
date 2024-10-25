package services

import (
	"errors"
	"log"
	"user-service/pkg/utils"
	"user-service/src/models"

	"github.com/google/uuid"
)

func (s *UserService) CreateUser(userDto *models.CreateUserDTO) (*models.User, error) {
	err := uuid.Validate(userDto.ID)
	if err != nil {
		return nil, errors.New("invalid UUID")
	}

	if !utils.ValidateEmailRequirements(userDto.Email) {
		return nil, errors.New("invalid email")
	}

	if !utils.ValidatePasswordRequirements(userDto.Password) {
		return nil, errors.New("invalid password")
	}

	hashedPassword, err := utils.HashPassword(userDto.Password)
	if err != nil {
		log.Printf("error hashing password: %v", err)
		return nil, errors.New("error hashing password")
	}

	user := models.User{
		ID:       userDto.ID,
		Email:    userDto.Email,
		Password: hashedPassword,
	}

	err = s.repo.Create(&user)
	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil, errors.New("error creating user")
	}

	return &user, nil
}
