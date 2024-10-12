package methods

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Task3() {
	circle := Circle{
		Radius: 5.5,
	}
	fmt.Printf("Circle Radius: %g\nArea: %.2f\n", circle.Radius, circle.Area())
}
