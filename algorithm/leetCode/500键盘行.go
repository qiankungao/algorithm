package leetCode

import "strings"

func FindWords(words []string) []string {

	res := make(map[int32]int, 26)
	for _, s := range "qwertyuiop" {
		res[s] = 1
	}
	for _, s := range "asdfghjkl" {
		res[s] = 2
	}
	for _, s := range "zxcvbnm" {
		res[s] = 4
	}

	str := make([]string, 0, len(words))
	for _, w := range words {
		size := 0
		ww:= strings.ToLower(w)
		for _, s := range ww {
			size += res[s]
		}
		if size == len(w) || size == len(w)*2 || size == len(w)*4 {
			str = append(str, w)
		}
	}
	return str
}
