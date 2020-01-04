package plate_test

import (
	"testing"

	"github.com/nahuelcaceres/quote-builder/plate"
	"github.com/stretchr/testify/assert"
)

func TestGetPrice(t *testing.T) {

	model := plate.NewPlate(15.50, 4, 3)

	service := plate.NewService()

	platesQuantity := float64(service.GetPlatesQuantity(6, 6, model))
	priceCalculated := service.GetPrice(6, 6, model)

	var expectedPrice float64 = platesQuantity * model.Price

	assert.Equal(t, expectedPrice, priceCalculated, "The two values should be the same.")

	//assert.Nil(t, nil)
}
