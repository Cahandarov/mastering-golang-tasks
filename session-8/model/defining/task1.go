package defining

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func Task1() {
	fmt.Println("Task 1  ****************")
	var circle Shape
	var rectangle Shape
	circle = Circle{Radius: 3}
	rectangle = Rectangle{Width: 9, Height: 6}
	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Rectangle area:", rectangle.Area())

}
