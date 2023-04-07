package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pedroluti1989/twiter/bd"
	"github.com/pedroluti1989/twiter/models"
)

/* Paara registrar un Usuario */
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Los datos del usuario son incorrectos"+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.EditarUsuario(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ucurrio un error al editar el usuario. Intente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado editar el usuario", 400)
		return
	}
	/* envio el header, para ve el mensaje 201 de exito */
	w.WriteHeader(http.StatusCreated)
}
