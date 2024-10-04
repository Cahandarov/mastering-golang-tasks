package main

import "fmt"

var a int = 10
var b int = 20

func swap(val1 *int, val2 *int) {
	var c int = 0
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
	c = *val1     //Passing val1 value to c
	*val1 = *val2 //Passing val2 value to val1
	*val2 = c     //Passing c value to val2
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
}
func task3() {
	swap(&a, &b)
}
