package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	printElements(array)
	sum := add(array)
	fmt.Printf("-   Sum: %d\n", sum)
	reverseArray(array)
}

func printElements(array [5]int) {
	fmt.Printf("-   First element: %d\n", array[0])
	fmt.Printf("-   Last element: %d\n", array[len(array)-1])
}

func add(array [5]int) int {
	var sum int = 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
func reverseArray(array [5]int) {

	fmt.Print("-   Reversed array: [")
	for j := len(array) - 1; j >= 0; j-- {
		if j > 0 {
			fmt.Printf("%d, ", array[j])
		} else {
			fmt.Printf("%d]\n", array[j])
		}
	}
}
