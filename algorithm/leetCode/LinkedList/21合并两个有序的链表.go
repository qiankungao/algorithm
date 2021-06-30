package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	HeadNode *ListNode
}

//判断链表是否空
func (l *List) IsEmpty() bool {
	if l.HeadNode == nil {
		return true
	}
	return false
}

//长度
func (l *List) Length() int {
	cur := l.HeadNode
	count := 0
	if cur.Next != nil {
		count++
		cur = cur.Next
	}
	return count
}

//从链表尾部添加元素
func (l *List) Append(data int) {
	node := &ListNode{Val: data}
	if l.IsEmpty() {
		l.HeadNode = node
	} else {
		cur := l.HeadNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 在链表指定位置插入元素
func (l *List) Insert(index, data int) {
	if index > l.Length() { // 大于长度，就在尾部插入
		l.Append(data)
	} else {
		pre := l.HeadNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		node := &ListNode{Val: data}
		next := pre.Next
		pre.Next = node
		node.Next = next
	}
}

//遍历
func (l *List) ShowList() {
	if l.IsEmpty() {
		return
	}
	cur := l.HeadNode
	for cur.Next != nil {
		fmt.Printf("%d ->", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
func main() {
	list := List{}
	list.Append(1)
	list.Append(4)
	list.Append(6)
	list.Append(3)
	fmt.Println("长度：", list.Length())
	list.ShowList()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
