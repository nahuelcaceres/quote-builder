package placa_test

import (
	"testing"

	"github.com/nahuelcaceres/quote-builder/placa"
	"github.com/stretchr/testify/assert"
)

func TestCalcularPrecio(t *testing.T) {

	model := placa.NuevaPlaca(15.50, 1.20, 2.40, 1.25)

	servicio := placa.NuevoServicio()

	cantidadPlacas := float64(servicio.CalcularCantidadPlacas(6, 6, model))
	precioCalculado := servicio.CalcularPrecio(6, 6, model)

	var precioEsperado float64 = cantidadPlacas * model.Precio

	assert.Equal(t, precioEsperado, precioCalculado, "The two values should be the same.")

	//assert.Nil(t, nil)
}
