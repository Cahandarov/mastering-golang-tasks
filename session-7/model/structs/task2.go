package structs

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func Task2() {
	rec := Rectangle{Width: 10.5, Height: 20.5}
	fmt.Printf("Width: %g, Height: %g\nArea: %g\nPerimeter: %g\n", rec.Width, rec.Height, rec.Area(), rec.Perimeter())
}
