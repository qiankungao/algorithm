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
	result := leetCode.Intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(result)

	//str := make(map[string]int, 0)
	//str["name"] = 1
	//str["age"] = 2
	//_, ok := str["name"]
	//fmt.Println("ok:", ok)
	//
	//for k := range str {
	//	fmt.Println("k:", k)
	//}

}
