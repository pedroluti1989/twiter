package routers

import (
	"net/http"
	"twiter/bd"
	"twiter/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UserID = IDUsuario
	relacion.UserRelacionID = ID

	status, err := bd.InsertarRelacion(relacion)
	if err != nil {
		http.Error(w, "ocurrio un error al insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
