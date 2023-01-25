package routers



import (
	"strings"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pedroluti1989/twiter/models"
	"github.com/pedroluti1989/twiter/bd"
)
/*Email valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string
/* Procesar el token */
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
    miClave := []byte("Pedro_desarrollo_go_2214")
	claims := &models.Claim{}
    
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de Token Invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}