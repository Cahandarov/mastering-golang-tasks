package benchmarking

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return arr
}

//func Task5() {
//	fmt.Println("Task-5    **************8")
//	numbers := []int{64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90}
//	fmt.Println(BubbleSort(numbers))
//}
