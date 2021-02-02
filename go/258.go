package leetcode

func addDigits(num int) int {
	for {
		var sum int
		for _, v := range strconv.Itoa(num) {
			n, _ := strconv.Atoi(string(v))
			sum += n
		}
		if sum < 10 {
			return sum
		}
		num = sum
	}
}
