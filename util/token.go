package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const signKey = "ex2022"

type TokenRaw struct {
	Uid uint
	jwt.StandardClaims
}

func GenToken(uid uint) string {
	tokenRaw := TokenRaw{
		uid,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 10,
			Issuer:    "ex",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenRaw)
	tokenStr, err := token.SignedString([]byte(signKey))
	if err != nil {
		panic(err)
	}
	return tokenStr
}

func ParseToken(tokenStr string) (*TokenRaw, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenRaw{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenRaw); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
