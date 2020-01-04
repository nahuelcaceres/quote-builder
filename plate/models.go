package plate

// Placa modelo
type Plate struct {
	Price  float64
	Height float64
	Width  float64
}

func NewPlate(price float64, height float64, width float64) Plate {
	return Plate{
		Price:  price,
		Height: height,
		Width:  width,
	}
}
