package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"museum-api/database"
	"museum-api/models"
	"net/http"
	"strings"
	"time"
)

// GenerateToken generates a JWT com o email enviado como parâmetro, e retorna o token gerado e um possível erro
// Lembre de mudar o tempo ta com 72h andré seu animal
// Lembrar de mudar o Secret do JWT pra um que não seja uma merda
func GenerateToken(email string) (string, error) {
	secret := viper.GetString("JWT_SECRET")

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := viper.GetString("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	return token, err
}

func ValidateTokenMiddleware(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		context.Abort()
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		context.Abort()
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token email claim"})
		context.Abort()
		return
	}

	// Aqui você deve obter o ID do gerente do banco de dados usando o e-mail do token
	managerID, err := GetManagerIDByEmail(email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get manager ID"})
		context.Abort()
		return
	}

	// Define o ID do gerente no contexto Gin
	context.Set("manager_id", managerID)
	context.Next()
}

func GetManagerIDByEmail(email string) (uint, error) {
	var managerID uint
	err := database.DB.Model(&models.Manager{}).Select("id").Where("email = ?", email).First(&managerID).Error
	if err != nil {
		return 0, err
	}
	return managerID, nil
}
