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
