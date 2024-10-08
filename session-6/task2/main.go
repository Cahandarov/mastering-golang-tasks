package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := array[:]
	subSlice := slice[2:6]
	fmt.Printf("Sub-slice: %d \n", subSlice)
	appendSlice := append(subSlice, 10, 11, 12)
	fmt.Printf("Appended slice: %d \n", appendSlice)
	fmt.Printf("Original array length, capacity and referance: %d, %d, %d\n", len(array), cap(array), &array[2])
	fmt.Printf("subSlice slice length, capacity and referance: %d, %d, %d\n", len(subSlice), cap(subSlice), &subSlice[0])
	fmt.Printf("appendSlice slice length, capacity and referance: %d, %d, %d\n", len(appendSlice), cap(appendSlice), &appendSlice[0])
}
