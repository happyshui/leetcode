package leetcode

func isHappy(n int) bool {
	nmap := make(map[int]int)
	var sum int
	for {
		for _, v := range strconv.Itoa(n) {
			num, _ := strconv.Atoi(string(v))
			square := num * num
			sum += square
			_, ok := nmap[sum]
			if ok {
				return false
			}
		}
		if sum == 1 {
			return true
		}
		nmap[sum] = 1
		n = sum
		sum = 0
	}
}
