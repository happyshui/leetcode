package leetcode

func missingNumber(nums []int) int {
	length := len(nums)
	var lsum int
	var nsum int
	lsum = (length + 1) * length / 2
	for i := 0; i < length; i++ {
		nsum += nums[i]
	}
	return lsum - nsum

}

// Test site: bit operation
func missingNumber(nums []int) int {
	length := len(nums)
	var xor int
	for i := 1; i <= length; i++ {
		xor = xor ^ i
	}
	for i := 0; i < length; i++ {
		xor = xor ^ nums[i]
	}
	return xor
}
