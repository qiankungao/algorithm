// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/leetCode/dp"
	"fmt"
)

func main() {

	//ops := "(){[]}" 0 1 1 2 3
	fmt.Println(dp.MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))

}
