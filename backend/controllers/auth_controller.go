package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Denol007/social-network-prototype/backend/models"
	"github.com/Denol007/social-network-prototype/backend/services"
	"github.com/gin-gonic/gin"
)

// Структура для входных данных регистрации
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Обработчик регистрации нового пользователя
func RegisterUser(c *gin.Context) {
	var req RegisterRequest

	// Валидация входных данных
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
	}

	if err := services.RegisterUser(&user, req.Password); err != nil {
		log.Printf("Ошибка регистрации пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser обрабатывает запрос логина пользователя
func LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	// Вызов сервиса для проверки учетных данных
	user, token, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	// Успешный вход — возвращаем успешный ответ с токеном
	c.SetCookie(
		"auth_token",
		token,
		int(24*time.Hour.Seconds()), // Время жизни cookie, совпадающее с временем жизни JWT (24 часа)
		"/",
		"",   // Домен; оставьте пустым, чтобы использовался домен запроса
		true, // Secure: если сервер HTTPS, установите в true
		true, // HttpOnly, чтобы токен не был доступен через JavaScript
	)

	// Возвращаем успешный ответ без токена, так как он уже в cookies
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
