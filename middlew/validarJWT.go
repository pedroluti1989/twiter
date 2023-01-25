package middlew

import (
	"net/http"
	"github.com/pedroluti1989/twiter/routers"
)

func ValidarJWT ( next http.HandlerFunc) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){
		_,_,_,err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w,"No es un Token valido"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w,r)
	}
}