package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	secretKey string
}

func NewAuthSerive(secretKey string) *AuthService {
	return &AuthService{secretKey: secretKey}
}

func (authService *AuthService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(authService.secretKey))
}

func (authService *AuthService) ValidateToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(authService.secretKey), nil
	})

	if err != nil {
		return "", nil
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if userID, ok := claims["userID"].(string); ok {
			return userID, nil
		}
		return "", errors.New("userID não encontrado")
	}

	return "", errors.New("token invalido")
}
