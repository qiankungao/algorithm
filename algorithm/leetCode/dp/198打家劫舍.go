package dp

/*
1,2,3,1
2,1,1,2
*/

//第一版 有点low
func Rob1(nums []int) int {
	dp := make([]int, len(nums)+1)
	if len(nums) == 0 {
		return 0
	}
	dp[0], dp[1], dp[2] = 0, nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i+1] = max(dp[i-1], dp[i-2]) + nums[i]
	}
	if dp[len(dp)-1] >= dp[len(dp)-2] {
		return dp[len(dp)-1]
	}
	return dp[len(dp)-2]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//第k间的房子的对大值 = max(k-1的最大值，k-2+nums[i])
func Rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

//func Rob(nums []int) int {
//
//}
