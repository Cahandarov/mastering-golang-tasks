package embedding

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}
type Car struct {
	Vehicle
	FuelType string
}

func Task6() {
	car := Car{
		Vehicle: Vehicle{
			Make:  "Kia",
			Model: "K5",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}

	fmt.Printf("Make: %s\nModel: %s\nYear: %d\nFuel Type: %s", car.Make, car.Model, car.Year, car.FuelType)

}
