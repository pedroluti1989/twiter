package bd

import (
	"context"
	"log"
	"time"

	"twiter/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* para insrtar un usuario en la bd */
func InsertarUsuario(u models.Usuario) (string, bool, error) {

	log.Println(u)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	coleccion := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	registro := bson.M{
		"nombre":    u.Nombre,
		"apellido":  u.Apellido,
		"fechaNac":  u.FechaNac,
		"email":     u.Email,
		"password":  u.Password,
		"biografia": u.Biografia,
		"sitioWeb":  u.SitioWeb,
		"avatar":    u.Avatar,
		"banner":    u.Banner,
		"ubicacion": u.Ubicacion,
	}

	result, err := coleccion.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	ObjId := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil

}
