package utils

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("truffle_key")

type Claims struct {
	Hash string
	jwt.StandardClaims
}

func ReleaseToken(hash string, now int64) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Hash: hash,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  now,
			Issuer:    "org.truffle",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

func AuthJWT(token, name string) bool {
	if len(token) == 0 || !strings.HasPrefix(token, "TRUFFLE") {
		return false
	}
	token = token[7:]
	jwt, cliams, err := ParseToken(token)
	if err != nil || !jwt.Valid {
		return false
	}
	raw := DecodeAESWithKey(ToString(cliams.IssuedAt)+"TRUFFLE", cliams.Hash)
	if raw == name {
		return true
	}
	return false
}
