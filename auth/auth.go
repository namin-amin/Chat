package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims
//
// Auth jwt token schema
type JwtCustomClaims struct {
	Name  string `json:"name"` //name of the user
	Key   string `json:"key"`
	Admin bool   `json:"admin"` //is user admin
	jwt.RegisteredClaims
}

// NewToken
//
// Creates and returns new auth token
func NewToken(name string, isAdmin bool, key string) (string, error) {
	// https://echo.labstack.com/cookbook/jwt/
	claims := &JwtCustomClaims{
		Name:  name,
		Key:   key,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			//TODO add more details
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret")) //TODO secret key take from env
	if err != nil {
		return "", err
	}

	return t, err
}
