package main

import "fmt"

func main() {
	// 1 5    10 15 20
	//      6
	s := []int{1, 5, 10, 15, 20}
	fmt.Println(searchInsert(s, 5))
}
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
