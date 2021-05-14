package leetcode

// 暴力求解
func longestPalindrome(s string) string {
	length := len(s)
	if length == 1 {
		return s
	}
	var result string
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			str := s[i : j+1]
			if isPalindrome(str) && len(str) > len(result) {
				result = str
			}
		}
	}
	if result == "" {
		return s[0:1]
	}
	return result
}

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
