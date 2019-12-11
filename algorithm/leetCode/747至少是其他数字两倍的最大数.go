package leetCode

/*
   最大的数 至少是其他数的两倍 如果是，则返回最大元素的索引，否则返回-1。
	示例 1:

	输入: nums = [3, 6, 1, 0]
	输出: 1
	解释: 6是最大的整数, 对于数组中的其他整数,
	6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.
	示例 2:
	输入: nums = [1, 2, 3, 4]
	输出: -1
	解释: 4没有超过3的两倍大, 所以我们返回 -1.

*/

func DominantIndex(nums []int) int {
	//3, 6, 1, 0
	max := nums[0]
	pre := -1
	for i := 0; i < len(nums); i++ {
		if max <= nums[i] {
			max = nums[i]
			pre = i
		}
	}

	for _, v := range nums {
		if max < v*2 && max != v {
			return -1
			break
		}
	}
	return pre
}
