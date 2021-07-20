package leetcode

func titleToNumber(columnTitle string) int {
	length := len(columnTitle)
	ans := 0
	smap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
	for i := 0; i < length; i++ {
		tmp := power(length - 1 - i)
		ans += smap[string(columnTitle[i])] * tmp
	}
	return ans
}

func power(n int) int {
	p := 26
	res := 1
	for i := 0; i < n; i++ {
		res = res * p
	}

	return res
}
