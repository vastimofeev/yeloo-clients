package database

import (
	"fmt"
	"yeloo-clients/config"
	"yeloo-clients/internal/models"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

// BuildDSN формирует строку подключения к базе данных
func BuildDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode,
	)
}

func AutoMigrate(db *gorm.DB) {
	// Выполняем AutoMigrate
	err := db.AutoMigrate(&models.Profile{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("AutoMigrate completed successfully")
}

func Connect(dsn string) (*gorm.DB, error) {
	// Подключение к базе данных
	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database:", err)
		return nil, err
	}

	log.Info("Database connected successfully")

	// Выполняем AutoMigrate
	AutoMigrate(db)

	return db, nil
}
