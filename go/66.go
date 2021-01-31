package leetcode

func plusOne(digits []int) []int {
	idx := len(digits) - 1
	for i := idx; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
