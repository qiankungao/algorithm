package dp

/*
	示例 1:
	输入: cost = [10, 15, 20]
	输出: 15
	解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。

    示例 2:
	输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
	输出: 6
	解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。

*/

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < n; i++ {
		if dp[i-1] <= dp[i-2] {
			dp[i] = dp[i-1] + cost[i]
		} else {
			dp[i] = dp[i-2] + cost[i]
		}

	}
	if dp[n-1] >= dp[n-2] {
		return dp[n-2]
	}
	return dp[n-1]
}
