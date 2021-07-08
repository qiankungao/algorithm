package leetCode

import "testing"

func BenchmarkFindMinDifference(b *testing.B) {
	timePoints := []string{"23:59", "00:00"}
	for i := 0; i < b.N; i++ {
		FindMinDifference(timePoints)
	}
}
