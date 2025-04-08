package database

import (
	"fmt"
	"os"
	"os/exec"
	"yeloo-clients/config"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

// BuildDSN формирует строку подключения к базе данных
func BuildDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
}

func Connect(dsn string) (*gorm.DB, error) {
	// Подключение к базе данных
	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database:", err)
		return nil, err
	}

	log.Info("Database connected successfully")
	return db, nil
}

func RunMigrations(dsn string) error {
	cmd := exec.Command("atlas", "migrate", "apply", "--url", dsn)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
