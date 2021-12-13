package types

import (
	"graphql/resolvers"

	"github.com/graphql-go/graphql"
)

var NovedadType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Novedad",
	Fields: graphql.Fields{
		"localidad": &graphql.Field{Type: graphql.String},
		"fecha":     &graphql.Field{Type: graphql.String},
	},
})

var EnvioType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Envio",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"destinatario":  &graphql.Field{Type: graphql.String},
		"numedoDeEnvio": &graphql.Field{Type: graphql.String},
		"estado":        &graphql.Field{Type: graphql.String},
		"novedades": &graphql.Field{
			Type:    graphql.NewList(NovedadType),
			Resolve: resolvers.NovedadesResolver,
		},
	},
})

/*
	Aqui es donde defino los querys que se van a poder ejecutar.
	Se hace un nuevo objeto de GraphQL con la siguiente informacion.

	Name  : -> Siempre le dejo Query, nunca lo cambio y es para que sepa cual es el root de los queryS.
	Fields: -> En este caso indico cada uno de los querys que se pueden ejecutar ( "envio" y "envios" en mi caso )

	La estructura de cada "field" es la siguiente:
	  - Type:    -> Tipo de objeto que va a devolver. Estan en este mismo file y se describe propiedad por propiedad
	  - Args:    -> Parametros que se reciben desde la query, por ejemplo en este caso voy a recibir el id del dato que quiero mostrar
	  - Resolve: -> Fucion que resuelve la consulta. Debe devolver un objeto ( interface ) con la misma estructura descripta en EnvioType
	                para el caso de envio. Para ver este resolver anda a /resolvers/resolvers.go
*/
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"envio": &graphql.Field{
			Type: EnvioType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.EnvioResolver,
		},
		"envios": &graphql.Field{
			Type: graphql.NewList(EnvioType),
			Args: graphql.FieldConfigArgument{
				"texto": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.EnviosResolver,
		},
	},
})
