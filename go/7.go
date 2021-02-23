package leetcode

func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		fmt.Println(tmp)
		x = x / 10
		fmt.Println(x)
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}

	return tmp
}
