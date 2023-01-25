package routers

import (
	"encoding/json"
	"net/http"
	"github.com/pedroluti1989/twiter/bd"
	"github.com/pedroluti1989/twiter/models"
)

/* Paara registrar un Usuario */
func VerPerfil(w http.ResponseWriter, r *http.Request){
	var t  models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	var ID = r.URL.Query().Get("id")
	if len(ID)< 1 {
		http.Error(w, "Debe enviar un Id valido", http.StatusBadRequest)
		return
	}

	perfil, err := bd.GetUsuario(ID)

	if err != nil {
		http.Error(w, "Ucurrio un erro al recuperar el usuario"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}