package models

//Password  representa o modelo de recuperacao de senha
type Password struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
