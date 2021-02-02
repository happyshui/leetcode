package leetcode

// 1. 计算速度过慢
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var count int
	for i := 2; i < n; i++ {
		if isprime(i) {
			count++
		}
	}
	return count
}

func isprime(num int) bool {
	for j := 2; j*j <= num; j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}
