package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//判断链表是否空
func (l *ListNode) IsEmpty() bool {
	if l.Length() == 0 {
		return true
	}
	return false
}

//长度
func (l *ListNode) Length() int {
	cur := l
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//从链表尾部添加元素
func (l *ListNode) Append(data int) {
	node := &ListNode{Val: data}
	if l.IsEmpty() {
		l = node
	} else {
		cur := l
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 在链表指定位置插入元素
func (l *ListNode) Insert(index, data int) {
	if index > l.Length() { // 大于长度，就在尾部插入
		l.Append(data)
	} else {
		pre := l
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
func (l *ListNode) ShowList() {
	if l.IsEmpty() {
		return
	}
	cur := l
	for cur != nil {
		fmt.Printf("%d ->", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

//删除指定元素的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Val == val {
		head = head.Next
		if head == nil {
			break
		}
		cur = head

	}
	for cur != nil {
		next := cur.Next
		if next == nil {
			break
		}
		if next.Val == val {
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l1 := &ListNode{Val: head.Val}
	head = head.Next
	for head != nil {
		node := head
		head = head.Next
		node.Next = l1
		l1 = node

	}
	return l1
}

func main() {
	list := &ListNode{Val: 1}
	list.Append(2)
	list.Append(3)
	list.Append(3)
	list.Append(21)
	list.Append(1)
	//list.ShowList()
	//reversePrint(list)
	fmt.Println(isPalindrome(list))
}

/*
	输入：head = [1,3,2]
	输出：[2,3,1]
*/
//从尾到头打印链表
func reversePrint(head *ListNode) []int {
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	s := make([]int, count)
	cur = head
	for cur != nil {
		count--
		s[count] = cur.Val
		cur = cur.Next
	}
	return s
}

/*
	1 -> 3 -> 4 -> 6
                5 -> 7-> 9
*/
//合并两个有序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := new(ListNode)
	pre := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}

	if l2 != nil {
		pre.Next = l2
	}
	return prehead.Next
}

// 合并两个链表
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	return nil
}

func isPalindrome(head *ListNode) bool {
	sli := []int{}
	for ; head != nil; head = head.Next {
		sli = append(sli, head.Val)
	}
	// 0,1,2,3,4,5
	for i, v := range sli[:len(sli)/2] {
		if v != sli[len(sli)-i-1] {
			return false
		}
	}
	return true
}
