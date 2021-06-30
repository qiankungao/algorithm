package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{1, 4}, {2, 3}, {5, 10}}
	fmt.Println(merge(nums))
}

/*
	1    3
      2      6
			5     10
						15   18


*/
func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
	假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。


*/
func findDuplicate(nums []int) int {
	//1,3,4,2,2
	flag := make([]int, len(nums))
	for _, v := range nums {
		flag[v-1]++
		if flag[v-1] == 2 {
			return v
		}
	}
	return 0
}
