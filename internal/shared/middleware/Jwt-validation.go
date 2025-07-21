package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

// ValidateToken valida si un token JWT es válido
func ValidateToken(tokenString string) error {
	if tokenString == "" {
		return errors.New("token no proporcionado")
	}

	// Remover el prefijo "Bearer " si está presente
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	// Parsear y validar el token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea el esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma no válido")
		}
		return jwtKey, nil
	})

	if err != nil {
		return errors.New("token inválido: " + err.Error())
	}

	if !token.Valid {
		return errors.New("token no válido")
	}

	return nil
}

// IsTokenValid verifica si un token es válido (retorna boolean)
func IsTokenValid(tokenString string) bool {
	return ValidateToken(tokenString) == nil
}

// ExtractTokenFromHeader extrae el token del header Authorization
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("header Authorization vacío")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("formato de token inválido, debe comenzar con 'Bearer '")
	}

	token := authHeader[len("Bearer "):]
	if token == "" {
		return "", errors.New("token vacío después del prefijo Bearer")
	}

	return token, nil
}

// JWTAuthMiddleware middleware para Gin que valida JWT y controla acceso a recursos
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Extraer token del header
		tokenString, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Acceso denegado para el recurso solicitado",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		// Validar token
		if err := ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Acceso denegado para el recurso solicitado",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		// Token válido, permitir acceso al recurso
		c.Next()
	}
}

// ValidateResourceAccess valida si un token permite acceso a un recurso específico
func ValidateResourceAccess(authHeader string) (bool, error) {
	// Extraer token
	tokenString, err := ExtractTokenFromHeader(authHeader)
	if err != nil {
		return false, err
	}

	// Validar token
	if err := ValidateToken(tokenString); err != nil {
		return false, err
	}

	return true, nil
}
