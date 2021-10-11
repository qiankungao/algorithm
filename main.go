package main

import "fmt"

//var a []int
//var c []int // 第三者

func main() {

	b := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("b:=", b[1])
	f(b)
	fmt.Println("b:=", b[1])
}

func f(b []int) []int {
	a := b[:2]

	// 新的切片 append 导致切片扩容
	var c []int
	c = append(c, b[:2]...)
	c[1] = 999
	fmt.Printf("a: %p\nb: %p\nc: %p\n", &a[0], &b[0], &c[0])

	return c
}
