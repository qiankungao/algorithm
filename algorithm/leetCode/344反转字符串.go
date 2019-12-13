package leetCode

import "fmt"

/*
1,2,3,4,5,6,7
*/
func ReverseString(s []byte) {
	fmt.Println(s)
	first, end := 0, len(s)-1
	for {
		if first >= end {
			break
		}
		s[first], s[end] = s[end], s[first]
		first++
		end--
	}
	fmt.Println(s)
}
