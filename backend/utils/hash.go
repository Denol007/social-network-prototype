package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Функция для хеширования пароля
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("ошибка при хешировании пароля: %v", err)
	}
	return string(hashedPassword), nil
}

// Функция для сравнения пароля с хешем
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("неверный пароль: %v", err)
	}
	return nil
}
