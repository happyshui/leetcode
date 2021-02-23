package leetcode

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	var norepeat int
	for i := range s {
		smap := make(map[byte]bool)
		var tmp int
		for j := i; j < length; j++ {
			_, ok := smap[s[j]]
			if ok {
				tmp = j - i
				break
			} else {
				smap[s[j]] = true
			}
			tmp = j + 1 - i
		}
		if norepeat < tmp {
			norepeat = tmp
		}
	}
	return norepeat
}

// sliding window algorithm
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	var norepeat int
	smap := make(map[byte]int)
	right, norepeat := -1, 0
	for i := 0; i < length; i++ {
		var cur int
		if i != 0 {
			delete(smap, s[i-1])
		}
		for right+1 < length && smap[s[right+1]] == 0 {
			smap[s[right+1]]++
			right++
		}
		cur = right - i + 1
		if cur > norepeat {
			norepeat = cur
		}
	}
	return norepeat
}
