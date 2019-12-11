package dp

/*
从左上角到右下角的最小路径和
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
注意边界的处理
*/
func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0]) //行、列
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] <= dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] += dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
