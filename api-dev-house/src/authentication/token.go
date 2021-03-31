package authentication

import (
	"api-dev-house/src/config"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateToken ... retorna um token de acesso ao sistema
func GenerateToken(userId int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // Time to expire
	permissions["userID"] = userId
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

//ValidateToken ... verifica se o token passado na requisição é válido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificateKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""

}

func getVerificateKey(token *jwt.Token) (interface{}, error) {
	if _, status := token.Method.(*jwt.SigningMethodHMAC); !status {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil

}

//ExtractUserId ... retorna o id do usuario no token
func ExtractUserId(r *http.Request) (int64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificateKey)

	if err != nil {
		return int64(0), err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["userID"]), 10, 64)
		if err != nil {
			return int64(0), err
		}
		return userID, nil
	}

	return int64(0), errors.New("token inválido")

}
