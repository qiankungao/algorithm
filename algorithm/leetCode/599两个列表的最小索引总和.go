package leetCode

func FindRestaurant(list1 []string, list2 []string) []string {
	//
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}
	m := make(map[string]int, len(list1))
	for index, v := range list1 {
		m[v] = index
	}
	res := make([]string, 0)
	min := len(list1) + len(list2)
	for index, v := range list2 {
		if val, ok := m[v]; ok {
			if min > index+val {
				res = make([]string, 0)
				res = append(res, v)
				min = index + val
			} else if min == index+val {
				res = append(res, v)
			}
		}
	}
	return res
}
