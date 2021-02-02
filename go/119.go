package leetcode

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+2)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res[rowIndex]
}

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+2)
	var result []int
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
		if i == rowIndex {
			result = res[i]
			break
		}
	}
	return result
}
