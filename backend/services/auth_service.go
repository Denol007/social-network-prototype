package services

import (
	"fmt"

	"github.com/Denol007/social-network-prototype/backend/models"
	"github.com/Denol007/social-network-prototype/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

// Функция для регистрации нового пользователя
func RegisterUser(user *models.User, password string) error {
	// Хешируем пароль перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка при хешировании пароля: %v", err)
	}
	user.PasswordHash = string(hashedPassword)

	// Сохраняем пользователя в базе данных
	err = repository.CreateUser(user)
	if err != nil {
		return fmt.Errorf("ошибка сохранения пользователя в базе данных: %v", err)
	}
	return nil
}
