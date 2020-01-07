package placa

import "math"

type Servicio struct {
}

func NuevoServicio() Servicio {
	return Servicio{}
}

func (servicio Servicio) CalcularPrecio(alto float64, ancho float64, placa Placa) float64 {
	var cantidadDePlacas = float64(servicio.CalcularCantidadPlacas(alto, ancho, placa))

	var precio = cantidadDePlacas * placa.Precio

	return precio
}

func (servicio Servicio) CalcularCantidadPlacas(alto float64, ancho float64, placa Placa) int {
	var m2ACubrir = alto * ancho

	var cantidadDePlacas = int(math.Round(m2ACubrir / (placa.Alto * placa.Ancho)))

	return cantidadDePlacas
}
