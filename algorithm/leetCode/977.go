package leetCode

import "fmt"

/*
输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/
func SortedSqSuares(A []int) []int {

	res := make([]int, len(A))
	first, end, index := 0, len(A)-1, len(A)-1
	for {
		one := A[first] * A[first]
		two := A[end] * A[end]
		fmt.Println(one, two)
		if one >= two {
			res[index] = one
			first++
		} else {
			res[index] = two
			end--
		}
		index--
		if first > end {
			break
		}
	}
	return res
}
