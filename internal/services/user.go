package services

import (
	"fmt"
	"net/http"

	"github.com/iufb/goth/internal/models"
	"github.com/iufb/goth/internal/repositories"
	"github.com/iufb/goth/pkg/utils"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(payload *models.AuthPayload) (int, error) {
	u, err := s.repo.FindByEmail(payload.Email)
	if u != nil {
		return http.StatusUnprocessableEntity, fmt.Errorf("User already exists.")
	}
	hashed, err := utils.HashPass(payload.Password)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Server Error.")
	}
	err = s.repo.Create(models.User{
		Email:    payload.Email,
		Password: hashed,
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user error.")
	}

	return http.StatusCreated, nil
}

func (s *UserService) LoginUser(payload *models.AuthPayload) (int, error) {
	u, err := s.repo.FindByEmail(payload.Email)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("User not found.")
	}
	// check password
	if !utils.ValidatePassword(u.Password, payload.Password) {
		return http.StatusUnauthorized, fmt.Errorf("Wrong password, try again.")
	}
	return http.StatusOK, nil
}
