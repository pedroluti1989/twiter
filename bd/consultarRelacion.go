package bd

import (
	"context"
	"time"
	"twiter/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuarios */
func ConsultarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"userid":         t.UserID,
		"userrelacionid": t.UserRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
