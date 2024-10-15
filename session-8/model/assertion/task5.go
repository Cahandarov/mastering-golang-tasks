package assertion

import "fmt"

type EmptyInterface any

func Task5() {
	fmt.Println("Task 5  ****************")
	var i EmptyInterface
	i = 42
	val1, _ := i.(int)
	fmt.Printf("Value is of type int: %d\n", val1)
	i = "GoLang"
	val2, _ := i.(string)
	fmt.Printf("Value is of type string: %s\n", val2)
	i = 3.1415
	val3, _ := i.(float64)
	fmt.Printf("Value is of type float64: %f\n", val3)
}
