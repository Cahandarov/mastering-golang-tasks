package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("ad1d0c552b5f050c8a015a3fb587ebcd23c083910eb94d63a481e6887ce9c9c2")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Create(username string) (string, error) {

	expireTime := time.Now().Add(6 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)

}

func Verify(token string) (string, error) {

	parsedToken := token[len("Bearer "):]

	claims, err := jwt.ParseWithClaims(parsedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	return claims.Claims.(*Claims).Username, nil

}
