package leetcode

// Method of exhaustion: use sort
func longestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	var result int
	var tmp int
	sort.Ints(nums)
	for i := 0; i < length-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1]-nums[i] == 1 {
			tmp = tmp + 1
		} else {
			tmp = 0
		}
		if tmp > result {
			result = tmp
		}
	}
	return result + 1
}
