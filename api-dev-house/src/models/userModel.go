package models

import (
	"api-dev-house/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//User .., dados de um usuario
type User struct {
	Id        int64     `json:"id, omitempty"`
	Name      string    `json:"name,omitempty"`
	Login     string    `json:"login,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

//Prepare ... chama métodos para validar e formatar dados do usuário a ser inserido.
func (u *User) Prepare(isCreateUser bool) error {
	if err := u.validate(isCreateUser); err != nil {
		return err
	}

	if err := u.format(isCreateUser); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(isCreateUser bool) error {

	if u.Name == "" {
		return errors.New("Campo Nome é obrigatório")
	}
	if u.Login == "" {
		return errors.New("Campo Login é obrigatório")
	}
	if u.Email == "" {
		return errors.New("Campo Email é obrigatório")
	} else if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Email informado é inválido")
	}

	if isCreateUser && u.Password == "" {
		return errors.New("Campo Senha é obrigatório")
	}

	return nil
}

func (u *User) format(isCreateUser bool) error {

	u.Name = strings.TrimSpace(u.Name)
	u.Login = strings.TrimSpace(u.Login)
	u.Email = strings.TrimSpace(u.Email)

	if isCreateUser {
		passwordHash, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(passwordHash)

	}
	return nil

}
