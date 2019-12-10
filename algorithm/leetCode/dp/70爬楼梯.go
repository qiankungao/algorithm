package dp

/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。

	输入： 2
	输出： 2
	解释： 有两种方法可以爬到楼顶。
	1.  1 阶 + 1 阶
	2.  2 阶

	输入： 3
	输出： 3
	解释： 有三种方法可以爬到楼顶。
	1.  1 阶 + 1 阶 + 1 阶
	2.  1 阶 + 2 阶
	3.  2 阶 + 1 阶


*/
//dp解法
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//执行用时 :0 ms, 在所有 golang 提交中击败了100.00% 的用户
//内存消耗 :1.9 MB, 在所有 golang 提交中击败了96.96%的用户
func ClimbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first, two, ans := 1, 2, 0
	for i := 3; i <= n; i++ {
		ans = first + two
		first, two = two, ans
	}
	return ans
}
