package leetcode

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)
	result := 20000
	for i := 0; i < length; i++ {
		left := i + 1
		right := length - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			tmp := sum - target
			dtmp := result - target
			if abs(tmp) < abs(dtmp) {
				result = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
