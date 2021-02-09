package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	nmap := make(map[int]int)
	for i := range nums1 {
		nmap[nums1[i]] += 1
	}
	for i := range nums2 {
		v := nmap[nums2[i]]
		if v > 0 {
			result = append(result, nums2[i])
			nmap[nums2[i]]--
		}

	}
	return result
}
