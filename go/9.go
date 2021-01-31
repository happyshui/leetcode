package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversalnum := 0
	tmpnum := x
	for tmpnum > 0 {
		reversalnum = reversalnum*10 + tmpnum%10
		tmpnum = tmpnum / 10
	}
	return reversalnum == x

}
