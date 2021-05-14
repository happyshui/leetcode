package leetcode

func findDuplicate(nums []int) int {
	nmap := make(map[int]int)
	length := len(nums)
	var result int
	for i := 0; i < length; i++ {
		_, ok := nmap[nums[i]]
		if ok {
			nmap[nums[i]]++
		} else {
			nmap[nums[i]] = 1
		}
	}

	for k, v := range nmap {
		if v > 1 {
			result = k
		}
	}
	return result
}
