package leetcode

func findTheDifference(s string, t string) byte {
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	smap := make(map[string]int)
	tmap := make(map[string]int)
	for i := range s {
		_, ok := smap[string(s[i])]
		if ok {
			smap[string(s[i])] += 1
		} else {
			smap[string(s[i])] = 1
		}
	}
	for i := range t {
		_, ok := tmap[string(t[i])]
		if ok {
			tmap[string(t[i])] += 1
		} else {
			tmap[string(t[i])] = 1
		}
	}
	var result []byte
	for i := range letter {
		if smap[letter[i]] != tmap[letter[i]] {
			result = []byte(letter[i])
		}
	}
	return result[0]
}
