package leetcode

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	var shift uint32 = 31
	for num != 0 {
		result += (num & 1) << shift
		num = num >> 1
		shift -= 1
	}
	return result
}
