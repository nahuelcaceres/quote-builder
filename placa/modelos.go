package placa

// Placa modelo
type Placa struct {
	Precio      float64
	Alto        float64
	Ancho       float64
	Profundidad float64
}

// NuevaPlaca factory method
func NuevaPlaca(precio float64, alto float64, ancho float64, profundidad float64) Placa {
	return Placa{
		Precio:      precio,
		Alto:        alto,
		Ancho:       ancho,
		Profundidad: profundidad,
	}
}
