package leetCode

/*
	f(1)=0,f(2)=1
	F(n) = F(n-1)+F(n-2)
*/
func Fib(N int) int {
	if N == 0 {
		return 0
	}
	first, two := 0, 1
	ans := two
	for i := 2; i <= N; i++ {
		ans = first + two
		first, two = two, ans
	}
	return ans
}
