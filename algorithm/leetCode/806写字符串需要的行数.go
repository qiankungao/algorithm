package leetCode

func NumberOfLines(widths []int, s string) []int {
	line := 1
	count := 0
	for index := range s {
		if count+widths[s[index]-97] > 100 {
			count = widths[s[index]-97]
			line++
		} else {
			count += widths[s[index]-97]
		}
	}
	return []int{line, count}
}
