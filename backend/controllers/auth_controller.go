package controllers

import (
	"log"
	"net/http"

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
