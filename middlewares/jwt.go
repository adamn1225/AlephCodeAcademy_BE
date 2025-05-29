// jwt.go
package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	Role string `json:"role"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(role string, id uint) (string, error) {
	claims := AuthClaims{
		Role: role,
		ID:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
