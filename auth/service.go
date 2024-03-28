package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type JwtService struct {
}

func NewService() *JwtService {
	return &JwtService{}
}

var SECRET_KEY = []byte("secret-key")

func (s *JwtService) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
