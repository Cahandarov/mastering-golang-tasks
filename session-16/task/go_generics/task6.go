package go_generics

import "fmt"

func MinMax[T int | float64](slice []T) (T, T) {
	n := len(slice) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice[0], slice[n]
}
func Task6() {
	fmt.Println("Task 6   *******************")

	minV, maxV := MinMax([]int{10, 20, 5, 8, 15})
	fmt.Println("Min: ", minV, ", Max :", maxV)

}
