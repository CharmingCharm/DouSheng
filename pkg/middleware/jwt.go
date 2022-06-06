package middleware

import (
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

func GenToken(username string, id int64) (string, error) {
	claims := Claims{
		Username: username,
		Id:       id,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	if err != nil {
		return constants.DefaultErrString, status.TokenErr
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, status.TokenErr
	}

	if tokenClaims == nil {
		return nil, status.TokenErr
	}

	claims, ok := tokenClaims.Claims.(*Claims)

	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, status.TokenErr
}
