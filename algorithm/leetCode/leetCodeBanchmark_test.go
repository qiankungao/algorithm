package leetCode

import "testing"

func BenchmarkFindMinDifference(b *testing.B) {
	timePoints := []string{"23:59", "00:00"}
	for i := 0; i < b.N; i++ {
		FindMinDifference(timePoints)
	}
}

func BenchmarkNumberOfLines(b *testing.B) {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		NumberOfLines(widths, s)
	}
}

func BenchmarkUniqueMorseRepresentations(b *testing.B) {
	words := []string{"gin", "zen", "gig", "msg"}
	for i := 0; i < b.N; i++ {
		UniqueMorseRepresentations(words)
	}
}
