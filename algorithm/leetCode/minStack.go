package leetCode

import "fmt"

/*
	实现一个栈，带有出栈（pop），入栈（push），取最小元素（getMin）三个方法。要保证这三个方法的时间复杂度都是O（1）
*/
const (
	defaultSize = 100
)

type Stack struct {
	size    int   //容量
	top     int   //栈顶元素
	data    []int //存放元素
	minData []int //辅助数组，便于求最小值
}

func NewStack() *Stack {
	return &Stack{
		size:    defaultSize,
		top:     -1,
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

//判断是否为空
func (this *Stack) IsEmpty() bool {
	return this.top == -1
}

//判断栈是否已满
func (this *Stack) IsFull() bool {
	return this.size == len(this.data)
}

//清空栈
func (this *Stack) ClearStack() {
	this.top = -1
	this.data = make([]int, this.size)
	this.minData = make([]int, this.size)
}

//栈的长度
func (this *Stack) GetLength() int {
	return this.top + 1
}

//入栈
func (this *Stack) Push(p int) {
	if this.IsFull() {
		fmt.Println("the Stack is full")
		return
	}

	this.data = append(this.data, p)
	this.top += 1

	if len(this.minData) == 0 {
		this.minData = append(this.minData, p)
	} else {
		if p < this.minData[len(this.minData)-1] {
			this.minData = append(this.minData, p)
		}
	}

}

//出栈
func (this *Stack) Pop() int {
	if this.IsEmpty() {
		fmt.Println("the Stack is Empty")
		return 0
	}
	tmp := this.data[this.top]
	if tmp == this.minData[len(this.minData)-1] {
		this.minData = this.minData[:len(this.minData)-1]
	}
	this.top--
	return tmp
}

//遍历栈
func (this *Stack) Traverse() {
	if this.IsEmpty() {
		fmt.Println("the Stack is Empty")
		return
	}

	for i := this.top; i >= 0; i-- {
		fmt.Print(this.data[i], " ")
	}
}

//取栈的最小元素
func (this *Stack) GetMin() int {
	for i := this.top; i >= 0; i-- {
		if this.data[i] == this.minData[len(this.minData)-1] {
			return this.data[i]
		}
	}
	return 0
}
