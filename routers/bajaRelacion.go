package routers

import (
	"net/http"
	"twiter/bd"
	"twiter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UserID = IDUsuario
	relacion.UserRelacionID = ID

	status, err := bd.BorrarRelacion(relacion)
	if err != nil {
		http.Error(w, "ocurrio un error al borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
