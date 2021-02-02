package leetcode

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	var res strings.Builder
	var start int = 0
	for i := 1; i < len(s); i++ {
		if string(s[i]) != string(s[start]) {
			s1 := strconv.Itoa(i - start)
			res.WriteString(s1)
			res.WriteString(string(s[start]))
			start = i
		}
	}
	res.WriteString(strconv.Itoa(len(s) - start))
	res.WriteString(string(s[start]))
	return res.String()
}
