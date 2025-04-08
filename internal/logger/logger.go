package logger

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger initializes the logrus logger with predefined settings.
func InitLogger() {
	// Устанавливаем формат логов в формате JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Настраиваем ротацию логов
	logFile := &lumberjack.Logger{
		Filename:   "logs/app.log", // Путь к файлу логов
		MaxSize:    100,            // Максимальный размер файла в мегабайтах
		MaxBackups: 7,              // Максимальное количество резервных копий
		Compress:   true,           // Сжимать старые файлы логов
	}

	// Создаем MultiWriter для вывода в консоль и файл
	fileWriter := io.MultiWriter(logFile)
	consoleWriter := io.MultiWriter(os.Stdout)

	// Настраиваем уровни логирования
	log.SetOutput(io.MultiWriter(fileWriter, consoleWriter))
	log.SetLevel(log.DebugLevel)

	// Логируем запуск
	log.WithField("timestamp", time.Now().Format(time.RFC3339)).Info("Logger initialized")
}
