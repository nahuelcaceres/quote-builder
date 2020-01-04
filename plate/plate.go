package plate

import "math"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (service Service) GetPrice(height float64, width float64, plate Plate) float64 {
	var cantidadDePlacas = float64(service.GetPlatesQuantity(height, width, plate))

	var price = cantidadDePlacas * plate.Price

	return price
}

func (service Service) GetPlatesQuantity(height float64, width float64, plate Plate) int {
	var m2ACubrir = height * width

	var cantidadDePlacas = int(math.Round(m2ACubrir / (plate.Height * plate.Width)))

	return cantidadDePlacas
}
