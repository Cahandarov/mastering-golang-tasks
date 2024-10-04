package main

import "fmt"

var x int = 5

func incrementByValue(val int) {
	fmt.Printf("Value of x before incrementing by value:%d\n", x)
	val++
	fmt.Printf("Value of x after incrementing by value:%d\n", x)
}

func incrementByPointer(ptr *int) {
	fmt.Printf("Value of x before incrementing by pointer:%d\n", x)
	*ptr++
	fmt.Printf("Value of x after incrementing by pointer:%d\n", x)
}

func task2() {
	incrementByValue(x)
	incrementByPointer(&x)

}
