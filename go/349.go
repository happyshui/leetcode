package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	nmap := make(map[int]bool)
	for i := range nums1 {
		nmap[nums1[i]] = false
	}
	for i := range nums2 {
		_, ok := nmap[nums2[i]]
		if ok {
			nmap[nums2[i]] = true
		}
	}
	for k, v := range nmap {
		if v == true {
			result = append(result, k)
		}
	}
	return result
}
