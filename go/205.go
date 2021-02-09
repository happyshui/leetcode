package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[string]int)
	tmap := make(map[string]int)
	var news string
	var newt string
	for i := 0; i < len(s); i++ {
		sv, ok := smap[string(s[i])]
		if ok {
			news += "-" + strconv.Itoa(sv)
		} else {
			smap[string(s[i])] = i
			news += "-" + strconv.Itoa(i)
		}
	}
	for i := 0; i < len(t); i++ {
		tv, ok := tmap[string(t[i])]
		if ok {
			newt += "-" + strconv.Itoa(tv)

		} else {
			tmap[string(t[i])] = i
			newt += "-" + strconv.Itoa(i)
		}
	}
	return news == newt
}
