package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Denol007/social-network-prototype/backend/config"
	_ "github.com/lib/pq"
)

func NewDBConnection(cfg *config.Config) (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname, cfg.Database.Sslmode)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Соединение с базой данных установлено")
	return db, nil
}
