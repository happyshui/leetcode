package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	var strs string
	for i := range s {
		if isstrornum(s[i]) {
			strs += string(s[i])
		}
	}

	strs = strings.ToLower(strs)
	left := 0
	right := len(strs) - 1
	for left < right {
		if strs[left] != strs[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isstrornum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
