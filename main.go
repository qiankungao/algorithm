// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/leetCode"
	"fmt"
)

//["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
func main() {

	//ops := "(){[]}"
	result := leetCode.MissingNumber([]int{1, 0})
	fmt.Println(result)

}
