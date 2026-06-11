package middleware

import (
	"fmt"
	"os"
	"time"

	"example.com/go/backend/models"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken(userID int, email string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT"))
	fmt.Println(jwtSecret)
	claims := models.Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
