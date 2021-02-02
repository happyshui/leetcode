package leetcode

func moveZeroes(nums []int) {
	var count int
	for c := range nums {
		if nums[c] == 0 {
			count += 1
		}
	}
	times := len(nums) - 1
	for j := 0; j < len(nums)-count; j++ {
		for i := times; i > 0; i-- {
			if nums[i-1] == 0 {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
	}
}
