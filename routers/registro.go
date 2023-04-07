package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"twiter/bd"
	"twiter/models"
)

/* Paara registrar un Usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	log.Println("usuario desde router", t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email es un dato requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El Password debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "El usuario con ese email ya esta registrado", 400)
		return
	}

	_, status, err := bd.InsertarUsuario(t)
	if err != nil {
		http.Error(w, "Ucurrio un error al insertar el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar en usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
