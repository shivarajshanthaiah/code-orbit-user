package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string
	Email  string
	jwt.StandardClaims
}

func GenerateToken(key, email, userID string) (string, error) {
	expTime := time.Now().Add(time.Hour * 10).Unix()

	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
			Subject:   email,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", email, err.Error())
		return "", err
	}
	return signedToken, nil
}
