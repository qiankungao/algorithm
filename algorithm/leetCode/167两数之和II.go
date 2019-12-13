package leetCode

func TwoSum(numbers []int, target int) []int {
	first, end := 0, len(numbers)-1
	for numbers[first]+numbers[end] > target {
		if numbers[first]+numbers[end] > target {
			end--
		} else {
			first++
		}
	}
	return []int{first + 1, end + 1}
}
