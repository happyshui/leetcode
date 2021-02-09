package leetcode

func countSegments(s string) int {
	var count int
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			count += 1
		}
	}
	return count
}
