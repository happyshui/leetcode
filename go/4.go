package leetcode

// Method of exhaustion
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	length := len(nums1)
	if length == 0 {
		return result
	}
	if length == 1 {
		return float64(nums1[0])
	}
	mid := length / 2
	if length%2 == 0 {
		result = (float64(nums1[mid]) + float64(nums1[mid-1])) / 2
	} else {
		result = float64(nums1[mid])
	}
	return result
}
