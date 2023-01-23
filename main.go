package main

import (
	"log"
	"github.com/pedroluti1989/twiter/handlers"
	"github.com/pedroluti1989/twiter/bd"

)

func main (){
	if bd.ChequearConexion() == 0 {
		log.Fatal("Fallo la conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}