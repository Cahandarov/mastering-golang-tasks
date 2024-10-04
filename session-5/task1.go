package main

import "fmt"

func task1() {
	var x int = 10
	ptrX := &x
	fmt.Printf("Value of x: %d\nAddress of x: %p\nValue at pointer: %d\n", x, ptrX, *ptrX)
}
