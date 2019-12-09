package leetCode

//输入: [1,2,3,1]
//输出: true
//执行用时 :32 ms, 在所有 golang 提交中击败了31.30% 的用户
//内存消耗 :8.1 MB, 在所有 golang 提交中击败了48.37%的用户
func ContainsDuplicateVersion1(nums []int) bool {

	res := make(map[int]int)
	for _, v := range nums {
		res[v] += 1
	}
	for _, v := range res {
		if v > 1 {
			return true
			break
		}
	}
	return false
}
func ContainsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	min := nums[0]
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	res := make([]int, max-min+1)
	for _, v := range nums {
		res[v-min] += 1
		if res[v-min] > 1 {
			return true
			break
		}
	}
	return false
}
