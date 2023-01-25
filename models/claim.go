package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Para validar el token */
type Claim struct{
	Email string `json:"email"`
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	jwt.StandardClaims
}