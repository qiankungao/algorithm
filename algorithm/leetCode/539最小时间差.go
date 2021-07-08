package leetCode

import (
	"sort"
	"strconv"
	"strings"
)

func FindMinDifference(timePoints []string) int {
	timePoi := make([]int, 0, len(timePoints))
	for _, s := range timePoints {
		t := strings.Split(s, ":")
		hour, _ := strconv.Atoi(t[0])
		second, _ := strconv.Atoi(t[1])
		timePoi = append(timePoi, hour*60+second)
	}
	if len(timePoi) == 0 {
		return 0
	}
	sort.Ints(timePoi)
	diff := timePoi[0] + 1440 - timePoi[len(timePoi)-1]
	for i := len(timePoi) - 1; i > 0; i-- {
		if diff > timePoi[i]-timePoi[i-1] {
			diff = timePoi[i] - timePoi[i-1]
		}
	}
	return diff
}
