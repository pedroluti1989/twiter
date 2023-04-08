package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID  string             `bson:"userid, omitempty" json:"userid"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
