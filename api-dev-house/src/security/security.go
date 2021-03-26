package security

import "golang.org/x/crypto/bcrypt"

//Hash ... recebe string e retorna o hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

//ComparePasswords ... compare senha pura com a senha com hash
func ComparePasswords(passWordHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passWordHash), []byte(password))
}
