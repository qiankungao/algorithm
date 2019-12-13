package leetCode

import (
	"strconv"
	"strings"
)

/*
	逢2 进1
*/
func AddBinary(a string, b string) string {
	one := strings.Split(a, "")
	other := strings.Split(b, "")
	lenA, lenB := len(one), len(other)
	lone := make([]string, 0)   //较长的数组
	sother := make([]string, 0) //较短的数组
	if lenA < lenB {
		lone = other
		sub := make([]string, lenB-lenA)
		sub = append(sub, one...)
		sother = sub
	} else {
		lone = one
		sub := make([]string, lenA-lenB)
		sub = append(sub, other...)
		sother = sub
	}
	for index, v := range sother {
		if v == "" {
			sother[index] = "0"
		}
	}
	jin := 0
	res := make([]string, len(lone)+1)
	j := len(lone)
	for i := len(lone) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(lone[i])
		b, _ := strconv.Atoi(sother[i])
		add := a + b + jin
		if add == 2 {
			res[j] = "0"
			jin = 1
		}
		if add == 3 {
			res[j] = "1"
			jin = 1
		}
		if add == 1 {
			res[j] = "1"
			jin = 0
		}
		if add == 0 {
			res[j] = "0"
			jin = 0
		}
		j--
	}
	res[j] = strconv.Itoa(jin)
	if res[0] == "0" {
		return strings.Join(res[1:], "")
	}

	return strings.Join(res, "")
}
