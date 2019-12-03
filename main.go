// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/leetCode"
	"fmt"
)

//["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
func main() {

	var minStack = leetCode.Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)

	minStack.Pop()
	fmt.Println("最小值：", minStack.GetMin())

	minStack.Pop()
	fmt.Println("最小值：", minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	top := minStack.Top()
	fmt.Println("最小值：", top, minStack.GetMin())

	minStack.Pop()
	minStack.Pop()
	fmt.Println("最小值：", minStack.GetMin())

}
