package leetcode

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 1 && nums[0] == target {
		return []int{0, 0}
	}
	left := 0
	right := length - 1
	var start int = -1
	var end int = -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					start = i
				}
			}
			for j := mid; j <= right; j++ {
				if nums[j] == target {
					end = j
				}
			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{start, end}
}
