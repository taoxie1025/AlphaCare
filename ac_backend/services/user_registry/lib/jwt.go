package lib

import (
	"alphacare/ac_backend/services/user_registry/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func GenJWT(jwtKey string, expireTimeDelta time.Duration, userInfo *model.User) (string, error) {
	expirationTime := time.Now().Add(expireTimeDelta)
	claims := &Claims{
		Email: userInfo.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "User Registry Service",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
