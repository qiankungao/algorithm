package leetCode

/*
	泰波那契序列 Tn 定义如下：
	T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
	给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

	示例 1：
	输入：n = 4
	输出：4
	解释：
	T_3 = 0 + 1 + 1 = 2
	T_4 = 1 + 1 + 2 = 4

	示例 2：
	输入：n = 25
	输出：1389537

*/
//此方法超时
func TribonacciVersionOne(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-3)
}
func Tribonacci(n int) int {
	tri := make([]int, 38)
	tri[0] = 0
	tri[1] = 1
	tri[2] = 1

	for i := 3; i <= 37; i++ {
		tri[i] = tri[i-1] + tri[i-2] + tri[i-3]
	}
	return tri[n]
}
