package leetcode

func reverseWords(s string) string {
	result := []byte{}
	length := len(s)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}

		for end := start; end < i; end++ {
			result = append(result, s[i-end+start-1])
		}

		for i < length && s[i] == ' ' {
			i++
			result = append(result, ' ')
		}
	}
	return string(result)
}
