package models

import (
	"database/sql"
	"time"
)

// Модель пользователя
type User struct {
	ID           string         `json:"id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"-"`
	Bio          sql.NullString `json:"bio,omitempty"`
	AvatarURL    sql.NullString `json:"avatar_url,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
