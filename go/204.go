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

// 2. 数学方法
func countPrimes(n int) int {
    res := 0
    if n < 2 {
        return res
    }
    isprime := make([]bool, n)
    for i := range isprime {
        isprime[i] = true
    }
    for i := 2; i < n; i++ {
        if isprime[i] {
            res++
            for j := 2*i; j < n; j = j+i {
                isprime[j] = false
            }
        }
    }
    return res
}
