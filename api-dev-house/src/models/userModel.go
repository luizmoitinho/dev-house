package models

import (
	"errors"
	"strings"
	"time"
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
func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.format()
	return nil
}

func (u *User) validate() error {

	if u.Name == "" {
		return errors.New("Campo Nome é obrigatório")
	}
	if u.Login == "" {
		return errors.New("Campo Login é obrigatório")
	}
	if u.Email == "" {
		return errors.New("Campo Email é obrigatório")
	}
	if u.Password == "" {
		return errors.New("Campo Senha é obrigatório")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Login = strings.TrimSpace(u.Login)
	u.Email = strings.TrimSpace(u.Email)

}
