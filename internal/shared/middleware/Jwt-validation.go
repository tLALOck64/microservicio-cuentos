package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/tLALOck64/microservicio-cuentos/internal/shared/response"
	"google.golang.org/api/option"
)

var (
	jwtKey      []byte
	firebaseApp *firebase.App
	authClient  *auth.Client
)

type User struct {
	UserID             string `json:"userId"`
	Email              string `json:"email"`
	FirebaseUID        string `json:"firebaseUid,omitempty"`
	TipoAutenticacion  string `json:"tipoAutenticacion"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	
	initFirebase()
}

func initFirebase() {
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credentialsPath == "" {
		log.Println("FIREBASE_CREDENTIALS_PATH no configurado, Firebase no estará disponible")
		return
	}

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Error inicializando Firebase: %v", err)
		return
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("Error obteniendo cliente de autenticación Firebase: %v", err)
		return
	}

	firebaseApp = app
	authClient = client
	log.Println("Firebase Admin SDK inicializado correctamente")
}

func isJWTFormat(token string) bool {
	parts := strings.Split(token, ".")
	return len(parts) == 3
}

func validateJWTToken(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("token no proporcionado")
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma no válido")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, errors.New("token JWT inválido: " + err.Error())
	}

	if !token.Valid {
		return nil, errors.New("token JWT no válido")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("claims JWT inválidos")
	}

	user := &User{
		UserID:            claims.Subject,
		Email:             "", 
		TipoAutenticacion: "local",
	}

	return user, nil
}

func validateFirebaseToken(tokenString string) (*User, error) {
	if authClient == nil {
		return nil, errors.New("Firebase no está configurado")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	decodedToken, err := authClient.VerifyIDToken(ctx, tokenString)
	if err != nil {
		return nil, errors.New("token Firebase inválido: " + err.Error())
	}

	user := &User{
		UserID:            decodedToken.UID,
		Email:             decodedToken.Claims["email"].(string),
		FirebaseUID:       decodedToken.UID,
		TipoAutenticacion: "firebase",
	}

	if emailClaim, exists := decodedToken.Claims["email"]; exists && emailClaim != nil {
		if email, ok := emailClaim.(string); ok {
			user.Email = email
		}
	}

	return user, nil
}

func ValidateToken(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("token no proporcionado")
	}

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	if isJWTFormat(tokenString) {
		user, err := validateJWTToken(tokenString)
		if err == nil {
			return user, nil
		}
		log.Printf("JWT validation failed, trying Firebase: %v", err)
	}

	user, err := validateFirebaseToken(tokenString)
	if err != nil {
		return nil, errors.New("token inválido en ambos sistemas: " + err.Error())
	}

	return user, nil
}

func IsTokenValid(tokenString string) bool {
	_, err := ValidateToken(tokenString)
	return err == nil
}

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

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

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
		user, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Acceso denegado para el recurso solicitado",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}
		tokenString, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.Next()
			return
		}
		user, err := ValidateToken(tokenString)
		if err != nil {
			log.Printf("Token validation failed in optional middleware: %v", err)
			c.Next()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func GetUserFromContext(c *gin.Context) (*User, bool) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	user, ok := userInterface.(*User)
	return user, ok
}

func ValidateResourceAccess(authHeader string) (*User, error) {
	tokenString, err := ExtractTokenFromHeader(authHeader)
	if err != nil {
		return nil, err
	}
	user, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return user, nil
}