package constantes

import a "github.com/fochoac/go-model-ws/prueba"

func Ok() a.CodigoRetorno {
	return a.CodigoRetorno{Codigo: "1",
		Mensaje: "Ok"}
}
func NoData() a.CodigoRetorno {
	return a.CodigoRetorno{Codigo: "2",
		Mensaje: "No existe la información solicitada"}
}
func Error() a.CodigoRetorno {
	return a.CodigoRetorno{Codigo: "3",
		Mensaje: "Error inesperado, Itente nuevamente más tarde"}
}
