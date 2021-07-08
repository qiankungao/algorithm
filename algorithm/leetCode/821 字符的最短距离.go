package leetCode

/*
s = "loveleetcode", c = "e"
[3,2,1,0,1,0,0,1,2,2,1,0]
*/
func ShortestToChar(s string, c byte) []int {

	max := len(s) << 1
	left, right := -max, max
	ans := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		ans = append(ans, max)
	}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			left, ans[i] = i, 0
		} else {
			ans[i] = i - left
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			right = i
		} else {
			if ans[i] > right-i {
				ans[i] = right - i
			}
		}
	}
	return ans
}
