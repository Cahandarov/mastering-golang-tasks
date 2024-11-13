package benchmarking

import "testing"

var numbers1 = []int{64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90}
var numbers2 = []int{64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90, 64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90, 64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90, 64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90, 64, 34, 25, 120, 453, 658, 86904, 92738, 1, 435, 56, 567, 987, 222, 12, 22, 11, 90}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("Benchmark of first slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleSort(numbers1)
		}
	})

	b.Run("Benchmark of second slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BubbleSort(numbers2)
		}
	})

}
