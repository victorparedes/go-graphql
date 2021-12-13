package resolvers

import (
	"fmt"
	"graphql/datos"
	"graphql/models"
	"strings"

	"github.com/graphql-go/graphql"
)

/*
  Estos son datos falsos para no pegarle a una base de datos.
  En una segunda instancia por ahi le pegue a la swapi o a la de LOTR
*/
var datosNovedad = datos.ObtenerNovedades()
var datosEnvio = datos.ObtenerEnvios()

/*
	Cada resolver se ejecuta solo si se necesita, es decir, en el caso de "NovedadesResolver" si no pido las novedades en mi consulta
	este resolver no se ejecuta y por ende, si necesitara ir a una base de datos a buscar datos no se realiza la busqueda.
*/
func NovedadesResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("* Realizando la busqueda de novedades *")
	var resultado []models.Novedad

	source, _ := p.Source.(models.Envio)

	for _, item := range datosNovedad {
		if source.Id == item.IdEnvio {
			resultado = append(resultado, item)
		}
	}

	return resultado, nil
}

func EnviosResolver(p graphql.ResolveParams) (interface{}, error) {
	texto, _ := p.Args["texto"].(string)
	var resultado []models.Envio
	for _, item := range datosEnvio {
		if strings.Contains(item.Destinatario, texto) {
			resultado = append(resultado, item)
		}
	}

	return resultado, nil
}

/*
	Este resolver busca el ID en los envios, si lo encuentra lo devuelve.
*/
func EnvioResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)

	for _, item := range datosEnvio {
		if item.Id == id {
			return item, nil
		}
	}

	return nil, nil
}
