package sort

import (
	"../../../BigDataAlgorithm/common"
)

/*
	计数排序：一种非基于比较的排序
	局限: 1. 仅适用于整数的排序
          2. 数列的最小值和最大值的差值不是特别大

*/
//计数排序 空间复杂度o(n),时间复杂度o(n+k)
func CountingSort(array []int) (sortArray []int) {
	//Step:1 循环数组，得到最大最小值
	max := common.GetMax(array)
	min := common.GetMin(array)
	//Step:2 开辟一个长度max-min+1大小的数组
	countArray := make([]int, max-min+1)
	//Step:3 循环原数组，把每个元素插入到对应的统计数组中,(并不是真正的插入，而是对应的位子的count+1)插入位置是v-min
	for _, v := range array {
		countArray[v-min] += 1
	}
	//Step:4 遍历统计数组，输出个数为count的min+index即为排好序的数组
	sortArray = make([]int, max-min+1)
	j := 0
	for index, value := range countArray {
		for i := 0; i < value; i++ {
			sortArray[j] = index + min
			j++
		}
	}
	return
}

func CountingSort2() {
	//Step:1 找出
}
