package leetCode

import "strings"

/*
	输入：S = "ab#c", T = "ad#c"
	输出：true
	解释：S 和 T 都会变成 “ac”

	输入：S = "ab##", T = "c#d#"
	输出：true
	解释：S 和 T 都会变成 “”。

	输入：S = "a##c", T = "#a#c"
	输出：true
	解释：S 和 T 都会变成 “c”。
*/

func BackspaceCompare(S string, T string) bool {
	return backspace(S) == backspace(T)
}

func backspace(s string) string {
	tmp := make([]string, len(s))
	for _, v := range s {
		switch string(v) {
		case "#":
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			}
		default:
			tmp = append(tmp, string(v))
		}
	}
	return strings.Join(tmp, "")
}
