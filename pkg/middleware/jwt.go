package middleware

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

// 根据用户名生成jwt
func GenToken(username string, id int64) (string, error) {
	claims := Claims{
		Username: username,
		Id:       id,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims == nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Claims)

	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
