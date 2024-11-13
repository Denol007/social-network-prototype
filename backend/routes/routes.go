package routes

import (
	"github.com/Denol007/social-network-prototype/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Регистрация маршрутов для API
	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
	}
}
