package bd

import (
	"context"
	"time"
	"twiter/models"
)

/* Insertar la relacion */
func InsertarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twiter")
	coleccion := db.Collection("relacion")

	_, err := coleccion.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
