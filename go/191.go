package leetcode

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		num = num & (num - 1)
		count += 1
	}
	return count
}
