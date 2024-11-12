package main

import (
	"log"

	"github.com/Denol007/social-network-prototype/backend/config"
	"github.com/Denol007/social-network-prototype/backend/repository"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключаемся к базе данных
	db, err := repository.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close() // Закрываем соединение при завершении работы приложения

	// Здесь запускаем веб-сервер или настраиваем маршрутизацию
	log.Println("Сервер успешно запущен")
}
