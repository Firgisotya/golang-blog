package helpers

import (
	"fmt"
	"go-blog/api/configs"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ExpireTime = time.Hour * 3

func CreateToken(userId uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["email"] = email
	claims["exp"] = ExpireTime

	fmt.Println(configs.Config.SECRETKEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.SECRETKEY), nil
	})
}

func ExtractToken(tokenStr string) (string, error) {
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	}

	return "", err
}
