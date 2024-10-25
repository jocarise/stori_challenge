package services

import (
	"errors"
	"log"
	"os"
	"user-service/pkg/jwt"
	"user-service/pkg/utils"
	"user-service/src/models"
)

func (s *UserService) AuthUser(userDto *models.AuthUserDTO) (string, error) {
	if !utils.ValidateEmailRequirements(userDto.Email) {
		return "", errors.New("invalid email or password")
	}

	if !utils.ValidatePasswordRequirements(userDto.Password) {
		return "", errors.New("invalid email or password")
	}

	user, err := s.repo.GetByEmail(userDto.Email)
	if err != nil {
		log.Printf("Error fetching user by email: %v", err)
		return "", errors.New("invalid email or password")
	}

	if err := utils.ComparePasswords(user.Password, userDto.Password); err != nil {
		log.Printf("Error compering passwords: %v", err)
		return "", errors.New("invalid email or password")
	}

	token, err := jwt.GenerateJWT(user, []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("error generating JWT token: %v", err)
		return "", errors.New("invalid email or password")
	}

	return token, nil
}
