package leetcode

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		b = a & b << 1
		a = sum
	}
	return a
}
