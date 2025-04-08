package dependencies

import (
	"yeloo-clients/internal/controllers"
	"yeloo-clients/internal/repositories"
	"yeloo-clients/internal/services"

	"gorm.io/gorm"
)

type Dependencies struct {
	ProfileController *controllers.ProfileController
}

func InitializeDependencies(db *gorm.DB) *Dependencies {
	// Инициализация репозитория
	profileRepo := &repositories.ProfileRepository{DB: db}

	// Инициализация сервиса
	profileService := &services.ProfileService{Repo: profileRepo}

	// Инициализация контроллера
	profileController := &controllers.ProfileController{Service: profileService}

	return &Dependencies{
		ProfileController: profileController,
	}
}
