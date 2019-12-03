package leetCode

import "strings"

/*给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

*/

func RemoveDuplicates(S string) string {
	str := make([]string, 0)
	for _, s := range S {
		if len(str) == 0 {
			str = append(str, string(s))
			continue
		}
		if str[len(str)-1] == string(s) {
			str = str[:len(str)-1]
		} else {
			str = append(str, string(s))
		}
	}

	return strings.Join(str, "")
}
