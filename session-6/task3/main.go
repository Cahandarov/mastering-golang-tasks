package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Printf("Before: length- %d, capacity- %d\n", len(slice), cap(slice))
	slice = append(slice, 1, 2)
	fmt.Printf("First append: length- %d, capacity- %d\n", len(slice), cap(slice))
	slice = append(slice, 3, 4)
	fmt.Printf("Second append: length- %d, capacity- %d\n", len(slice), cap(slice))
}
