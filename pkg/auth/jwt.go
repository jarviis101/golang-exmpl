package auth

import "github.com/dgrijalva/jwt-go"

func getToken(name string) (string, error) {
	signKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
	})
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
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
