package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService interface {
	Get(payload *Payload) (string, error)
	Verify(token string) (*jwt.Token, error)
}

type Payload struct {
	Phone string `json:"phone"`
}

type Claims struct {
	Phone string
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}

func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: "s3kp3t",
	}
}

func (s *jwtService) Get(p *Payload) (string, error) {
	claims := &Claims{
		p.Phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) Verify(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodRSA); !isValid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
