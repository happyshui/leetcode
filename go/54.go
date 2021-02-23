package leetcode

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	result := []int{}
	up := 0
	right := n - 1
	left := 0
	down := m - 1
	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}

		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}

		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}

		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			result = append(result, matrix[i][left])
		}

		left++
		if left > right {
			break
		}
	}
	return result
}
