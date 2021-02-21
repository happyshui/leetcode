package leetcode

func productExceptSelf(nums []int) []int {
	length := len(nums)
	right, result := make([]int, length), make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < length; i++ {
		result[i] = result[i] * right[i]
	}
	return result
}
