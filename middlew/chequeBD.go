package middlew

import (
	"net/http"
	"github.com/pedroluti1989/twiter/bd"

)

func ChequeoBD ( next http.HandlerFunc) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){
		if bd.ChequearConexion() == 0{
			http.Error(w,"Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}