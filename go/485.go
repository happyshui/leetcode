package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	tmp := 0
	for i := range nums {
		if nums[i] == 1 {
			tmp += 1
		} else {
			tmp = 0
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
