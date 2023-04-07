package bd

import (
	"context"
	"time"
	"twiter/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	coleccion := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := coleccion.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
