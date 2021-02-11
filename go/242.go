package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[byte]int)
	for i := range s {
		smap[s[i]] += 1
	}
	for i := range t {
		smap[t[i]] -= 1
	}
	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}
