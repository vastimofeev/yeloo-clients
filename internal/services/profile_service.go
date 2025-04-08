package services

import (
	"yeloo-clients/internal/models"
	"yeloo-clients/internal/repositories"
)

type ProfileServiceInterface interface {
	GetProfile(id uint) (*models.Profile, error)
	CreateProfile(profile *models.Profile) error
	DeleteProfile(id uint) error
}

type ProfileService struct {
	Repo *repositories.ProfileRepository
}

func (s *ProfileService) GetProfile(id uint) (*models.Profile, error) {
	return s.Repo.GetByID(id)
}

func (s *ProfileService) CreateProfile(profile *models.Profile) error {
	return s.Repo.Create(profile)
}

func (s *ProfileService) DeleteProfile(id uint) error {
	return s.Repo.DeleteByID(id)
}
