package main

import "fmt"

type Student struct {
	age int
}

func main() {
	s := ConvShopCell()

	fmt.Println(s[1].age == 0)
}

func ConvShopCell() map[int]Student {
	s := make(map[int]Student, 0)
	s[1] = Student{10}
	s[2] = Student{11}
	s[3] = Student{13}
	s[4] = Student{103}
	return s
}
