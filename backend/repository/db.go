package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Denol007/social-network-prototype/backend/config"
	_ "github.com/lib/pq"
)

// Объявляем глобальную переменную db
var db *sql.DB

// NewDBConnection устанавливает подключение к базе данных и сохраняет его в глобальной переменной db
func NewDBConnection(cfg *config.Config) error {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname, cfg.Database.Sslmode)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("Соединение с базой данных установлено")
	return nil
}

// CloseDBConnection закрывает подключение к базе данных
func CloseDBConnection() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
