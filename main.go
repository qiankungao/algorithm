// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/leetCode/dp"
	"fmt"
)

func main() {

	//ops := "(){[]}" 0 1 1 2 3
	fmt.Println(dp.MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))

}
