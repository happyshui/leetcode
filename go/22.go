package leetcode

func generateParenthesis(n int) []string {
	var result []string
	var gen func(left, right, n int, res string)
	gen = func(left, right, n int, res string) {
		if left == n && right == n {
			result = append(result, res)
			return
		}
		if left < n {
			gen(left+1, right, n, res+"(")
		}
		if left > right && right < n {
			gen(left, right+1, n, res+")")
		}
	}
	gen(0, 0, n, "")
	return result
}
