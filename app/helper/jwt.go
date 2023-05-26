package helper

import (
	"depobangunan/app/environment"
	"depobangunan/app/intface"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func JwtSign(p *intface.CheckAccount) (string, error) {
	payload := jwt.MapClaims{
		"email": p.Email,
		"id":    p.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	environment.ExportEnv()
	keyJwt := os.Getenv("KEYJWT")

	secretKey := []byte(keyJwt)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "Error signing token:", err
	}

	return signedToken, nil

}


func DecryptJWT(tokenString string, secretKey []byte) (*intface.JwtClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &intface.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*intface.JwtClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}