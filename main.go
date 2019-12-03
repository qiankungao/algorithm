// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/sort"
	"fmt"
)

func main() {

	s := sort.NewStack()

	s.Push(9)
	s.Push(3)
	s.Push(5)
	s.Push(7)
	s.Push(1)
	s.Push(2)
	s.Traverse()
	s.Pop()
	fmt.Println()
	s.Traverse()
	fmt.Println()
	fmt.Println("最小值", s.GetMin())
	s.Pop()
	s.Pop()
	fmt.Println("最小值", s.GetMin())

}
