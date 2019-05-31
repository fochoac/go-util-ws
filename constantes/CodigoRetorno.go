package constantes

import "github.com/fochoac/go-model-ws/prueba"

func Ok() prueba.CodigoRetorno {
	return prueba.CodigoRetorno{Codigo: "1",
		Mensaje: "Ok"}
}
func NoData() prueba.CodigoRetorno {
	return prueba.CodigoRetorno{Codigo: "2",
		Mensaje: "No existe la información solicitada"}
}
func Error() prueba.CodigoRetorno {
	return prueba.CodigoRetorno{Codigo: "3",
		Mensaje: "Error inesperado, Itente nuevamente más tarde"}
}
