package leetCode

import "fmt"

/*
转转数组 把数组中的元素向右旋转k个元素
	**/
func Rotate(nums []int, k int) {
	length, count := len(nums), 0
	for start := length - 1; count < len(nums)-1; start-- {
		current := start
		obj := nums[current]
		for {
			index := (current + k) % length
			tmp := nums[index]
			nums[index], obj = obj, tmp
			current = index
			count++
			if current == start {
				break
			}
		}
	}
	fmt.Println(nums)
}
