package leetCode

import "fmt"

type MinStack struct {
	size    int   //容量
	top     int   //栈顶元素
	data    []int //存放元素
	minData []int //辅助数组，便于求最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		size:    defaultSize,
		top:     -1,
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.top += 1

	if len(this.minData) == 0 {
		this.minData = append(this.minData, x)
	} else {
		if x <= this.minData[len(this.minData)-1] {
			this.minData = append(this.minData, x)
		}
	}
	fmt.Print("minData:", this.minData)
}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
	}
	tmp := this.data[this.top]
	this.top--
	if tmp == this.minData[len(this.minData)-1] {
		this.minData = this.minData[:len(this.minData)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if this.top < 0 {
		return 0
	}
	return this.data[this.top]
}

func (this *MinStack) GetMin() int {
	for i := this.top; i >= 0; i-- {
		if this.data[i] == this.minData[len(this.minData)-1] {
			return this.data[i]
		}
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
