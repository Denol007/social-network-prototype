package repository

import (
	"fmt"

	"github.com/Denol007/social-network-prototype/backend/models"
)

// Создание пользователя в базе данных
func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3) RETURNING id, created_at`

	err := db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %v", err)
	}
	return nil
}
