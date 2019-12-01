// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/sort"
	"fmt"
)

func main() {

	var array = []int{91, 95, 90, 93, 95, 92, 98, 99, 99, 93}
	fmt.Println("sort before:", array)

	sortArray := sort.CountingSort(array)
	fmt.Println("sort after:", sortArray)
}
