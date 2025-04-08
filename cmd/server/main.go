package main

import (
	"yeloo-clients/config"
	_ "yeloo-clients/docs"
	"yeloo-clients/internal/database"
	"yeloo-clients/internal/dependencies"
	"yeloo-clients/internal/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()
	// Формирование строки подключения
	dsn := database.BuildDSN(cfg)

	// Подключение к базе данных
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Инициализация зависимостей
	deps := dependencies.InitializeDependencies(db)

	// Настройка маршрутов
	router := gin.Default()
	routes.SetupRoutes(router, deps.ProfileController)

	// Swagger-документация
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	log.Infof("Starting server on port %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}
