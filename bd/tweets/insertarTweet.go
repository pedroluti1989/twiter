package tweets

import (
	"context"
	"time"

	"twiter/bd"
	"twiter/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* para insrtar un usuario en la bd */
func InsertarTweet(tweet models.Tweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("twiter")
	coleccion := db.Collection("tweet")

	registro := bson.M{
		"userid":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}

	result, err := coleccion.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
