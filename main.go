package main

import (
	"graphql/types"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	/*
		Aqui se crea el schema, es en donde se definen todas las querys que voy utilizar.
		En mi caso uso "envio" y "envios".

		Para ver como se configura estas querys anda y mira la variable QueryType en types/tipos.go
	*/
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: types.QueryType,
	})

	/*
		Utilice Handler para la implementacion del servidor de GraphQL pero es a gusto de cada uno.
	*/
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	/*
		Utilizo la libreria http nativa de go para no sumarle mas complejidad pero aqui he visto
		muchos ejemplos que utilizan gin gonic para aprovechar la seguridad y el uso de middlewares
	*/
	http.Handle("/graphql", h)
	http.ListenAndServe(":5000", nil)
}

/*
	Si no queres usar handler como servidor de GrapqhQL y preferis intentar ir por otro componente te dejo ejemplos

	con Go Nativo https://stackoverflow.com/questions/54448994/graphql-go-define-field-type-as-object
	con Gin Gonic: https://ajdelgados.com/2020/05/07/golang-graphql-para-una-api-usando-postgresql/
*/
