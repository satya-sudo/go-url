package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a signed JWT with a userId and role.
func GenerateJWT(secret, userId, role string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userId,                        // subject = user id
		"role": role,                          // user role
		"iat":  time.Now().Unix(),             // issued at
		"exp":  time.Now().Add(expiry).Unix(), // expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT verifies a JWT and returns claims if valid.
func ValidateJWT(secret, tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
