package leetcode

func containsDuplicate(nums []int) bool {
	nmap := make(map[int]bool)
	for i := range nums {
		_, ok := nmap[nums[i]]
		if ok {
			return true
		} else {
			nmap[nums[i]] = false
		}
	}
	return false
}

// second way: 1. sorted 2. diff(i, i+1)
