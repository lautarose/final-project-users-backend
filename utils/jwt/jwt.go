package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey = []byte("secret") //must be a local environment variable.
)

func GenerateToken(userID int) (string, error) {
	// Define la fecha de expiración del token
	expirationTime := time.Now().Add(24 * time.Hour) // Ejemplo: expira en 24 horas

	// Crea los claims (datos) del token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Id:        fmt.Sprintf("%d", userID),
	}

	// Crea el token JWT utilizando la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
