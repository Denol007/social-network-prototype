package repository

import (
	"database/sql"
	"errors"
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

func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, bio, avatar_url, created_at FROM users WHERE email=$1`
	row := db.QueryRow(query, email)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Bio, &user.AvatarURL, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	return &user, nil
}
