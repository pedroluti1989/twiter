package bd

import (
	"context"
	"time"
	"github.com/pedroluti1989/twiter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

/* para insrtar un usuario en la bd */
func InsertarUsuario(u models.Usuario) (string, bool, error){
	
	ctx, cancel := context.WithTimeout( context.Background(), 15*time.Second)
	defer cancel()

	db        := MongoCN.Database("twiter")
	coleccion := db.Collection("usuarios")

	u.Password,_ = EncriptarPassword(u.Password)

	result, err := coleccion.InsertOne(ctx, u)
	if err != nil{
		return "", false, err
	}

	ObjId := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil

}