package main

import "fmt"

var n1 int = 1
var n2 int = 2
var n3 int = 3

func doubleVariadicElements(nums ...*int) {

	fmt.Printf("Before doubling: %d %d %d\n", n1, n2, n3)

	for _, num := range nums {
		*num = *num * 2
	}
	fmt.Printf("After doubling: %d %d %d\n", n1, n2, n3)
}

func task4() {
	doubleVariadicElements(&n1, &n2, &n3)
}
