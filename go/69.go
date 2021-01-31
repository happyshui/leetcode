package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	right := x
	left := 0
	var ans int
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
