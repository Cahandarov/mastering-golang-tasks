package main

import "fmt"

var a int = 10
var b int = 20

func swap(val1 *int, val2 *int) {
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
	*val1, *val2 = *val2, *val1
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
}
func task3() {
	swap(&a, &b)
}
