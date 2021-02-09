package leetcode

func hammingDistance(x int, y int) int {
	n := x ^ y
	var count int
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
