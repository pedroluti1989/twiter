package bd

import (
	"context"
	"fmt"
	"time"
	"twiter/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Busco un usuario en la BD */
func GetUsuario(ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twiter")
	coleccion := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := coleccion.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Usuario no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil

}
