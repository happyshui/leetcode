package leetcode

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	backtrack(digits, 0, "", &result)
	return result
}

var dmap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func backtrack(digits string, idx int, res string, result *[]string) {
	if idx == len(digits) {
		*result = append(*result, res)
	} else {
		digit := string(digits[idx])
		str := dmap[digit]
		for i := 0; i < len(str); i++ {
			backtrack(digits, idx+1, res+string(str[i]), result)
		}
	}
}
