package leetcode

import "sort"

// 作弊
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums)
	return nums[length-k]
}

// 大顶堆
