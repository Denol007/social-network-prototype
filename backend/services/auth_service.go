package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Denol007/social-network-prototype/backend/models"
	"github.com/Denol007/social-network-prototype/backend/repository"
	"github.com/Denol007/social-network-prototype/backend/utils" // Импортируем utils
)

// Функция для регистрации нового пользователя
func RegisterUser(user *models.User, password string) error {
	// Хешируем пароль перед сохранением
	hashedPassword, err := utils.HashPassword(password) // Используем хеширование из utils
	if err != nil {
		return fmt.Errorf("ошибка при хешировании пароля: %v", err)
	}
	user.PasswordHash = hashedPassword

	// Сохраняем пользователя в базе данных
	err = repository.CreateUser(user)
	if err != nil {
		return fmt.Errorf("ошибка сохранения пользователя в базе данных: %v", err)
	}
	return nil
}

// Функция для входа пользователя
func LoginUser(email, password string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		log.Println("Ошибка при получении пользователя:", err)
		return nil, errors.New("неверный email или пароль")
	}

	// Сравниваем введенный пароль с хешем
	err = utils.CheckPasswordHash(password, user.PasswordHash) // Используем проверку пароля из utils
	if err != nil {
		return nil, errors.New("неверный email или пароль")
	}

	return user, nil
}
