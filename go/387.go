package leetcode

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	smap := make(map[byte]int, 26)
	var site int
	for i := range s {
		smap[s[i]] += 1
	}
	for i := range s {
		v := smap[s[i]]
		if v == 1 {
			site = i
			break
		} else {
			site = -1
		}
	}
	return site
}
