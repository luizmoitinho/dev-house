package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateToken ... retorna um token de acesso ao sistema
func GenerateToken(userId int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // Time to expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("secret")) //secrete key
}
