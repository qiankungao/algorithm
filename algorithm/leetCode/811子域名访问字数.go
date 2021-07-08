package leetCode

import (
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains))
	for _, s := range cpdomains {
		a := strings.Fields(s)
		count, _ := strconv.Atoi(a[0])
		m[a[1]] += count
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				m[s[i+1:]] += count
			}
		}
	}
	var res []string
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}
