package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type OrganisorClaims struct {
	OrganisorID uint   `json:"organisor_id"`
	Email       string `json:"email"`
	IsAdmin     bool `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint, email string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &OrganisorClaims{
		OrganisorID: userID,
		Email:       email,
		IsAdmin:     isAdmin, // Include admin status in claims
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*OrganisorClaims, error) {
	claims := &OrganisorClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
