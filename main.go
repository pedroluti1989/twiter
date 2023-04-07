package main

import (
	"log"
	"twiter/bd"
	"twiter/handlers"
)

func main() {
	if bd.ChequearConexion() == 0 {
		log.Fatal("Fallo la conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
