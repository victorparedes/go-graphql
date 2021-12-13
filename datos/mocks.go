package datos

import "graphql/models"

func ObtenerEnvios() []models.Envio {
	return []models.Envio{
		{
			Id:            1,
			Destinatario:  "Homero Simpson",
			NumedoDeEnvio: "000000000000001",
			Estado:        "Entregado",
		},
		{
			Id:            2,
			Destinatario:  "Bart Simpson",
			NumedoDeEnvio: "000000000000002",
			Estado:        "Pendiente",
		},
	}
}

func ObtenerNovedades() []models.Novedad {
	return []models.Novedad{
		{
			IdEnvio:   1,
			Fecha:     "07/12/2021",
			Localidad: "Moreno",
		},
		{
			IdEnvio:   1,
			Fecha:     "06/12/2021",
			Localidad: "CABA",
		},
		{
			IdEnvio:   1,
			Fecha:     "04/12/2021",
			Localidad: "Cordoba",
		},
		{
			IdEnvio:   2,
			Fecha:     "06/12/2021",
			Localidad: "Moron",
		},
	}
}
