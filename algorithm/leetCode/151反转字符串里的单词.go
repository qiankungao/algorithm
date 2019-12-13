package leetCode

import (
	"strings"
)

func ReverseWords(s string) string {
	s = strings.TrimSpace(s)
	str := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		if v != "" {
			str = append(str, v)
		}
	}
	if len(str) == 0 {
		return ""
	}
	first, end := 0, len(str)-1
	for {
		str[first], str[end] = str[end], str[first]
		first++
		end--
		if first > end || first == end {
			break
		}
	}
	return strings.Join(str, " ")
}
