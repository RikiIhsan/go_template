package utils

import (
	"github.com/RikiIhsan/lib/env"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stoken, err := token.SignedString([]byte(env.Get("SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	return stoken
}
