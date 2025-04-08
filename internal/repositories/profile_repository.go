package repositories

import (
	"errors"
	"yeloo-clients/internal/models"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	DB *gorm.DB
}

func (r *ProfileRepository) GetByID(id uint) (*models.Profile, error) {
	var profile models.Profile
	if err := r.DB.First(&profile, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Create(profile *models.Profile) error {
	return r.DB.Create(profile).Error
}

func (r *ProfileRepository) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Profile{}, id).Error
}
