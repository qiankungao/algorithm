package leetCode

/*
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

0: 1
1: 1 1
2: 1 2 1
3: 1 3 3 1
4: 1 4 6 4 1
*/

func Generate(numRows int) [][]int {
	//n[i][j] = n[i-1][j] + n[i-1][j-1]

	res := make([][]int, numRows)
	if numRows == 0 {
		return res
	}
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}

	for i := 2; i <= numRows; i++ {
		tmp := make([]int, i)
		tmp[0] = 1
		for j := 1; j < i-1; j++ {
			tmp[j] = res[i-2][j] + res[i-2][j-1]
		}
		tmp[i-1] = 1
		res[i-1] = tmp
	}
	return res
}

func GetRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res = []int{1}
	for i := 1; i <= rowIndex; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		for j := 1; j <= i-1; j++ {
			tmp[j] = res[j] + res[j-1]
		}
		tmp[i] = 1
		res = tmp
	}
	return res
}
