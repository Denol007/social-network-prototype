package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Структура для хранения информации о пользователе в токене
type Claims struct {
	UserID string `json:"user_id"` // Изменяем тип на string для UUID
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Генерация JWT токена
func GenerateJWT(userID string, email string) (string, error) { // Изменяем тип параметра на string
	// Получаем секретный ключ из переменной окружения
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
		return "", nil
	}

	// Создаем экземпляр Claims с данными пользователя
	claims := Claims{
		UserID: userID, // Используем string (UUID)
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // токен истекает через 24 часа
			Issuer:    "myapp",                               // имя приложения
		},
	}

	// Создаем новый JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
