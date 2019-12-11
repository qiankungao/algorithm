package dp

/*

机器人只能在向右、向下走，问从左上角走到右下角有多少条路径
dp 三步骤 ： 1. 确定dp[]数组元素的含义
            2. 找出状态转义方程，就是递推公式
            3. 确定初始值
dp[i][j] = dp[i-1][j]+dp[i][j-1]
m 是列
n 是行
*/
func UniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = 1
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
