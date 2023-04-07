package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"twiter/bd/tweets"
	"twiter/models"
)

/* Paara registrar el twwet de un usuario */
func GuardarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.Tweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := tweets.InsertarTweet(registro)

	if err != nil {
		http.Error(w, "Error al guardar el twwet en ña base de datos"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Ha ocurrido un error al guardar el tweet", 400)
		return
	}
	/* envio el header, para ve el mensaje 201 de exito */
	w.WriteHeader(http.StatusCreated)
}
