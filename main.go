// LargeDataAlgorithm project main.go
package main

import (
	"../BigDataAlgorithm/algorithm/leetCode"
	"fmt"
)

//["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
//[[6],[6],[],[],[],[5],[],[],[],[],[],[]]
//[null,true,6,6,true,true,5,true,-1,false,false,false]
func main() {

	//ops := "(){[]}"
	myQueue := leetCode.ConstructorQueue(3)
	fmt.Println(myQueue.EnQueue(6))
	fmt.Println(myQueue.Rear())
	fmt.Println(myQueue.Rear())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.EnQueue(5))
	fmt.Println(myQueue.Rear())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.Front())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.DeQueue())
	fmt.Println(myQueue.DeQueue())

}
