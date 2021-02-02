package leetcode

func maxSubArray(nums []int) int {
	sumMax := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sumMax {
			sumMax = nums[i]
		}
	}
	return sumMax
}
