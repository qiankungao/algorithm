package common

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func GetMax(array []int) int {
	if len(array) == 0 {
		return 0
	}
	max := INT_MIN
	for _, v := range array {
		if max <= v {
			max = v
		}
	}

	return max
}

func GetMin(array []int) int {
	if len(array) == 0 {
		return 0
	}
	min := INT_MAX
	for _, v := range array {
		if min >= v {
			min = v
		}
	}
	return min
}
