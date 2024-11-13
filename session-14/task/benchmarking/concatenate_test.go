package benchmarking

import "testing"

func BenchmarkConcatenate(b *testing.B) {
	sliceOfStrings := []string{"I", "hope", "one", "day", "i", "will", "be", "developer", "in", "Pasha"}
	b.Run("Concatenate with plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConcatenateWithPlus(sliceOfStrings)
		}
	})
	b.Run("Concatenate with strings.join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConcatenateWithJoin(sliceOfStrings)
		}
	})
}
