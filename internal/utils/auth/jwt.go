package auth

import (
	"ecommerce/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.StandardClaims
}
func GenTokenJWT(payload *PayloadClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.TokenSecret))
}

func CreateTokenJWT(payload *PayloadClaims) (string, error) {
	
	timeEx := global.Config.JWT.TokenExpirationTime
	if timeEx == "" {
		timeEx = "1h"
	}
	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expiration)
	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "ecommerce",
			Subject:   payload.Subject,
		},
	})
	
}

func VerifyTokenJWT(token string) (*PayloadClaims, error) {
	payload, err := jwt.ParseWithClaims(token, &PayloadClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(global.Config.JWT.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return payload.Claims.(*PayloadClaims), nil
}



