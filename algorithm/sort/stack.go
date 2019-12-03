package sort

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
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

//判断栈是否已满
func (s *Stack) IsFull() bool {
	return s.size == len(s.data)
}

//清空栈
func (s *Stack) ClearStack() {
	s.top = -1
	s.data = make([]int, s.size)
	s.minData = make([]int, s.size)
}

//栈的长度
func (s *Stack) GetLength() int {
	return s.top + 1
}

//入栈
func (s *Stack) Push(p int) {
	if s.IsFull() {
		fmt.Println("the Stack is full")
		return
	}

	s.data = append(s.data, p)
	s.top += 1

	if len(s.minData) == 0 {
		s.minData = append(s.minData, p)
	} else {
		if p < s.minData[len(s.minData)-1] {
			s.minData = append(s.minData, p)
		}
	}

}

//出栈
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("the Stack is Empty")
		return 0
	}
	tmp := s.data[s.top]
	if tmp == s.minData[len(s.minData)-1] {
		s.minData = s.minData[:len(s.minData)-1]
	}
	s.top--
	return tmp
}

//遍历栈
func (s *Stack) Traverse() {
	if s.IsEmpty() {
		fmt.Println("the Stack is Empty")
		return
	}

	for i := s.top; i >= 0; i-- {
		fmt.Print(s.data[i], " ")
	}
}

//取栈的最小元素
func (s *Stack) GetMin() int {
	for i := s.top; i >= 0; i-- {
		if s.data[i] == s.minData[len(s.minData)-1] {
			return s.data[i]
		}
	}
	return 0
}
