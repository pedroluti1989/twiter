package bd

import (
	"twiter/models"

	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(email string, pass string) (models.Usuario, bool) {
	usuario, encontrado, _ := ExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(pass)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuario, false
	}

	return usuario, true
}
