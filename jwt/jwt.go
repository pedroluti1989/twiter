package jwt

import (
	"time"
	"twiter/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("Pedro_desarrollo_go_2214")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellido":  t.Apellido,
		"fechaNac":  t.FechaNac,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"avatar":    t.Avatar,
		"banner":    t.Banner,
		"sitioWeb":  t.SitioWeb,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(miClave)

	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
