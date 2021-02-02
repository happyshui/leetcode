package leetcode

func romanToInt(s string) int {
	roman := make(map[string]int)
	roman["I"] = 1
	roman["IV"] = 4
	roman["V"] = 5
	roman["IX"] = 9
	roman["X"] = 10
	roman["XL"] = 40
	roman["L"] = 50
	roman["XC"] = 90
	roman["C"] = 100
	roman["CD"] = 400
	roman["D"] = 500
	roman["CM"] = 900
	roman["M"] = 1000

	var num int
	var i int = 0
	for i < len(s) {
		if i+1 < len(s) && roman[s[i:i+2]] > 0 {
			num += roman[s[i:i+2]]
			i = i + 2
		} else {
			num += roman[s[i:i+1]]
			i = i + 1
		}
	}
	return num
}
