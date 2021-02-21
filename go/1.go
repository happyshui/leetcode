package leetcode

func twoSum(nums []int, target int) []int {
	var idx1 int
	var idx2 int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 1; j < length; j++ {
			if nums[i]+nums[j] == target && i != j {
				idx1 = i
				idx2 = j
			}
		}
	}
	return []int{idx1, idx2}
}
