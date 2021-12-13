package models

/*
	Estos modelos son solo para el proposito de tener algun dato de prueba.
*/

type Envio struct {
	Id            int    `json:"id"`
	Destinatario  string `json:"destinatario"`
	NumedoDeEnvio string `json:"numedoDeEnvio"`
	Estado        string `json:"estado"`
}

type Novedad struct {
	IdEnvio   int    `json:"id"`
	Fecha     string `json:"fecha"`
	Localidad string `json:"localidad"`
}
