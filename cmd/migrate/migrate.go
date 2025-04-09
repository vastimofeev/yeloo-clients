package main

import (
	"log"
	"os"
	"os/exec"
	"yeloo-clients/config"
	"yeloo-clients/internal/database"
	"yeloo-clients/internal/models"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	db, err := database.Connect(database.BuildDSN(cfg))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Выполняем AutoMigrate
	err = db.AutoMigrate(&models.Profile{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("AutoMigrate completed successfully")

	// Генерация файла миграции с помощью Atlas
	migrationName := "auto_migrate_changes"
	cmd := exec.Command(
		"atlas", "migrate", "diff", migrationName,
		"--dir", "file://migrations",
		"--dev-url", database.BuildDSN(cfg),
		"--to", "file://migrations/schema.hcl",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to generate migration file: %v", err)
	}
	log.Println("Migration file generated successfully")
}
