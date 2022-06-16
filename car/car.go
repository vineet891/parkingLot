package car

import (
	"strings"
)

type Car struct {
	Number string
	Color  string
}

func NewCar(number, color string) *Car {
	return &Car{
		Number: number,
		Color:  color,
	}
}

func (this *Car) IsEqual(car Car) bool {
	return (this.Number == car.Number) && (strings.ToLower(this.Color) == strings.ToLower(car.Color))
}
