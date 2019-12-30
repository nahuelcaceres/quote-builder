package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrice(t *testing.T) {

	plate := Plate{12, 5, 150}

	//var error int = 0
	var priceCalculated float32 = 0
	var expectedPrice float32 = 1

	priceCalculated, _ = getPrice(2, 3, plate)

	assert.Equal(t, priceCalculated, expectedPrice, "The two values should be the same.")

	//assert.Nil(t, nil)
}
