// Package testutils implements utils for testing of microservice cars-app.
package testutils

import (
	"github.com/oeboada/cars-app/dto"
)

// GetCarInstance returns an instanced dto.Car for testing
func GetCarInstance() dto.Car {
	return dto.Car{
		Id:       "",
		Make:     "Fiat",
		Model:    "Grande Punto",
		Package:  "SE",
		Color:    "White",
		Year:     2015,
		Category: "Hatchback",
		Mileage:  59999,
		Price:    1099900,
	}
}
