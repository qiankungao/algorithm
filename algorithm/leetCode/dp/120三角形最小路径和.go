package dp

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：
[
    [2],
    [3,4],
	[6,5,7],
    [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
*/

//执行用时 :4 ms, 在所有 golang 提交中击败了97.32% 的用户
//内存消耗 :3.1 MB, 在所有 golang 提交中击败了70.97%的用户
func MinimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	dp[len(triangle)-1] = triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		dp[i] = triangle[i]
		for j := 0; j < len(triangle[i]); j++ {
			if dp[i+1][j] <= dp[i+1][j+1] {
				dp[i][j] += dp[i+1][j]
			} else {
				dp[i][j] += dp[i+1][j+1]
			}
			//dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + dp[i][j]
		}
	}
	return dp[0][0]
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
