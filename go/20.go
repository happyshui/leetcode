package leetcode

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	dicts := make(map[byte]byte)
	stack := []byte{}
	dicts[')'] = '('
	dicts[']'] = '['
	dicts['}'] = '{'

	for i := 0; i < len(s); i++ {
		if dicts[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != dicts[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1] // pop
			}
		} else {
			stack = append(stack, s[i]) // push
		}
	}
	return len(stack) == 0
}
