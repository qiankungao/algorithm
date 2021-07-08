package leetCode

func MinDeletionSize(strs []string) int {
	count := 0
	for c := 0; c < len(strs[0]); c++ {
		pre := strs[0][c]
		for r := 1; r < len(strs); r++ {
			if strs[r][c] >= pre {
				pre = strs[r][c]
			} else {
				count++
				break
			}
		}
	}
	return count
}
