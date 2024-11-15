package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Denol007/social-network-prototype/backend/models"
	"github.com/Denol007/social-network-prototype/backend/repository"
	"github.com/Denol007/social-network-prototype/backend/services"
	"github.com/Denol007/social-network-prototype/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
)

var oauthConf *oauth2.Config
var oauthStateString string

func init() {
	// Загружаем переменные из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	oauthStateString = os.Getenv("OAUTH_STATE")
}

// GoogleLogin перенаправляет на страницу Google OAuth
func GoogleLogin(c *gin.Context) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback обрабатывает ответ от Google после аутентификации
func GoogleCallback(c *gin.Context) {
	// Получаем код авторизации от Google
	code := c.DefaultQuery("code", "")

	// Обмен кода на токен доступа
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Ошибка обмена кода на токен: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при авторизации"})
		return
	}

	// Создаем HTTP клиент с токеном
	client := oauthConf.Client(context.Background(), token)

	// Создаем сервис для получения данных пользователя
	service, err := oauth2api.New(client)
	if err != nil {
		log.Printf("Ошибка создания сервиса Google: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при авторизации"})
		return
	}

	// Получаем информацию о пользователе
	userinfo, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Printf("Ошибка получения данных пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при авторизации"})
		return
	}

	// Проверяем, существует ли пользователь с таким email в базе данных
	user, err := repository.GetUserByEmail(userinfo.Email)
	if err != nil {
		// Если пользователя нет в базе данных, создаем нового
		if err.Error() == "пользователь не найден" {
			// Создаем нового пользователя
			newUser := models.User{
				Username: userinfo.Name,
				Email:    userinfo.Email,
			}

			// Регистрируем нового пользователя без пароля (передаем пустую строку)
			if err := services.RegisterUser(&newUser, ""); err != nil {
				log.Printf("Ошибка регистрации нового пользователя: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании нового пользователя"})
				return
			}

			// Успешная регистрация нового пользователя
			user = &newUser
		} else {
			log.Printf("Ошибка поиска пользователя: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска пользователя"})
			return
		}
	}

	// Генерация JWT токена для аутентификации
	jwtToken, err := utils.GenerateJWT(userinfo.Id, userinfo.Email)
	if err != nil {
		log.Printf("Ошибка генерации JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	// Устанавливаем JWT токен в куки
	c.SetCookie("jwt", jwtToken, 3600, "/", "localhost", false, true) // Срок действия 1 час

	// Возвращаем успешный ответ с данными пользователя
	c.JSON(http.StatusOK, gin.H{"message": "Authorization successful", "user": user})
}
