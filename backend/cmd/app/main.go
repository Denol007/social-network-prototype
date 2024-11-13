package main

import (
	"log"
	"os"

	"github.com/Denol007/social-network-prototype/backend/config"
	"github.com/Denol007/social-network-prototype/backend/repository"
	"github.com/Denol007/social-network-prototype/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %v", err)
	}

	// Проверяем наличие обязательной переменной окружения для JWT
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatal("JWT_SECRET_KEY не задан в переменных окружения")
	}

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключаемся к базе данных
	if err := repository.NewDBConnection(cfg); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer repository.CloseDBConnection() // Закрываем соединение при завершении работы приложения

	// Создаём экземпляр Gin
	router := gin.Default()

	// Регистрируем маршруты
	routes.RegisterRoutes(router)

	// Запускаем сервер на localhost:8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	log.Println("Сервер успешно запущен на http://localhost:8080")
}
