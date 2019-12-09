package leetCode

import "math"

/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
	如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
	注意你不能在买入股票前卖出股票。
	输入: [7,1,5,3,6,4]
	输出: 5
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

	输入: [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
func MaxProfit(prices []int) int {
	//7,1,5,3,6,4
	min := prices[0]
	sub := 0

	for _, v := range prices {
		//if v > min { //多加一个比价大小，会多耗时8ms
		//	sub = Max(sub, v-min)
		//}
		if v-min > sub {
			sub = v - min
		}
		if min > v {
			min = v
		}
	}
	return sub
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit1(prices []int) int {
	profit, min := 0, math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}

//动态规划
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	dpMin := make([]int, n)
	dpMin[0] = prices[0]
	for i := 1; i < n; i++ {
		if dpMin[i-1] < prices[i] {
			dpMin[i] = dpMin[i-1]
		} else {
			dpMin[i] = prices[i]
		}
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if prices[i]-dpMin[i-1] > dp[i-1] {
			dp[i] = prices[i] - dpMin[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
