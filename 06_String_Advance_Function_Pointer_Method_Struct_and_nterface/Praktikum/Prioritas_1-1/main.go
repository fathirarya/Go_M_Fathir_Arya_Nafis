package main

import (
	"fmt"
)

type Car struct {
	fuelType string
	fuel     float64
}

func (car Car) distanceTraveled() string {
	var distance float64

	switch car.fuelType {
	case "pertamax turbo":
		distance = car.fuel / 3.3
	case "pertamax":
		distance = car.fuel / 2.2
	case "pertalite":
		distance = car.fuel / 1.5
	default:
		distance = 0
	}

	return fmt.Sprintf("%.2f Mil", distance)
}

func main() {
	car1 := Car{
		fuelType: "pertamax turbo",
		fuel:     5,
	}

	car2 := Car{
		fuelType: "pertalite",
		fuel:     10,
	}

	fmt.Println("Jarak tempuh", car1.fuelType, "dengan jumlah", car1.fuel, "L adalah", car1.distanceTraveled())
	fmt.Println("Jarak tempuh", car2.fuelType, "dengan jumlah", car2.fuel, "L adalah", car2.distanceTraveled())
}
