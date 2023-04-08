package handlers

import (
	"log"
	"net/http"
	"os"
	"twiter/middlew"
	"twiter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/guardarTweet", middlew.ChequeoBD(middlew.ValidarJWT(routers.GuardarTweet))).Methods("POST")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.AltaRelacion))).Methods("POST")

	router.HandleFunc("/verTweets", middlew.ChequeoBD(middlew.ValidarJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")

	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")

	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.BajaRelacion))).Methods("DELETE")

	router.HandleFunc("/consultarRelacion", middlew.ChequeoBD(middlew.ValidarJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/verTweetsSeguidores", middlew.ChequeoBD(middlew.ValidarJWT(routers.GetTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
