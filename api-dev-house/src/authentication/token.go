package authentication

import (
	"api-dev-house/src/config"
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateToken ... retorna um token de acesso ao sistema
func GenerateToken(userId int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // Time to expire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) //secret key
}

func generateSecretKey() string {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(key)

}
