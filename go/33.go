package leetcode

// Method of exhaustion
func search(nums []int, target int) int {
	length := len(nums)
	idx := -1
	for i := 0; i < length; i++ {
		if nums[i] == target {
			idx = i
			return idx
		}
	}
	return idx
}
