package services

import (
	"fmt"
	"sharing-gogin/models"
	"sharing-gogin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey=utils.GoDotEnv("SECRET_KEY");

type authCustomClaims struct {
	User models.User
	jwt.StandardClaims
}

func GenerateToken(user models.User) string {
	claims := &authCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

}

func DecodingToken(tokenString string) jwt.MapClaims{
	
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	
	return claims;
}