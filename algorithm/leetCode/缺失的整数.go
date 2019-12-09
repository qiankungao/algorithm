package leetCode

/*
	1. 题目：一个无序数组里有99个不重复正整数，范围从1到100，唯独缺少一个整数。如何找出这个缺失的整数？
	2. 题目扩展：一个无序数组里有若干个正整数，范围从1到100，其中99个整数都出现了偶数次，只有一个整数出现了奇数次（比如1,1,2,2,3,3,4,5,5），如何找到这个出现奇数次的整数？
	3. 题目第二次扩展：一个无序数组里有若干个正整数，范围从1到100，其中98个整数都出现了偶数次，只有两个整数出现了奇数次（比如1,1,2,2,3,4,5,5），如何找到这个出现奇数次的整数？

*/

/*
	解题思路：
	0: 异或的原则，相同为0，不同为1
   1. 0异或任何数都得

*/
//题目扩展：一个无序数组里有若干个正整数，范围从1到100，其中99个整数都出现了偶数次，只有一个整数出现了奇数次（比如1,1,2,2,3,3,4,5,5）
func SingleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

//缺失得数 输入: [9,6,4,2,3,5,7,0,1]
//输出: 8 注意几种特殊的情况
func MissingNumber(nums []int) int {

	if len(nums) == 1 && nums[0] == 1 {
		return 0
	}
	min := nums[0]
	max := nums[0]
	sum := 0

	for _, v := range nums {
		sum += v
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	sum2 := 0
	for i := min; i <= max; i++ {
		sum2 += i
	}

	if sum2 == sum && min == 0 {
		return max + 1
	}
	return sum2 - sum
}
