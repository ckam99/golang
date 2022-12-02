package security

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Generate new token
func GenerateToken(data map[string]any, key string) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range data {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

// verify token
func VerifyToken(encodedToken string, key string) (*jwt.Token, error) {
	tokenString, err := DecodeToken(encodedToken)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// decode token
func DecodeToken(bearerToken string) (string, error) {
	if bearerToken == "" {
		return "", fmt.Errorf("missing auhorization header")
	}
	header := strings.Fields(bearerToken)
	if len(header) < 2 {
		return "", fmt.Errorf("invalid auhorization header format")
	}
	if strings.ToLower(header[0]) != "bearer" {
		return "", fmt.Errorf("unsupported auhorization type: %s", header[0])
	}
	return header[1], nil
}
