package defining

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func Task1() {
	fmt.Println("Task 1  ****************")
	var circle Shape
	var rectangle Shape
	circle = Circle{radius: 3}
	rectangle = Rectangle{width: 9, height: 6}
	fmt.Println("Circle area:", circle.area())
	fmt.Println("Rectangle area:", rectangle.area())
}
