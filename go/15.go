package leetcode

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := length - 1
		for left < right {
			res := []int{}
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, nums[i], nums[left], nums[right])
				result = append(result, res)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
