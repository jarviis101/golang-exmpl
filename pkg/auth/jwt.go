package auth

import (
	"github.com/dgrijalva/jwt-go"
	storage "prj/pkg/db/redis"
	"strings"
)

type Payload struct {
	Phone string `json:"phone"`
}

func getToken(p *Payload) (string, error) {
	token, _ := storage.GetTokenFromStorage(strings.Replace(p.Phone, "+", "ss:", 1))
	if token != "" {
		return token, nil
	}

	hash := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": p,
	})
	token, err := hash.SignedString([]byte("s3kp3t"))

	storage.SetTokenToStorage(strings.Replace(p.Phone, "+", "ss:", 1), token, 450)
	return token, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
