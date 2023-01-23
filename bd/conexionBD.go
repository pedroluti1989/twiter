package bd


import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

var MongoCN = ConectarBD()
var clientoptions = options.Client().ApplyURI("mongodb+srv://pedro89:3FyDzgOdQbEKXAJT@cluster0.ctwmf.mongodb.net/?retryWrites=true&w=majority")

/* Funcion q noes permite conexctar a la base de datos */
func ConectarBD() * mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientoptions)
	if err != nil{
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
		return client
	}
	log.Println("Conexion exitosa de BD")
	return client
}

/* Chekea la Conexion */

func ChequearConexion() int{
    err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 	1

}