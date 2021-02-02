package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	nmap := make(map[int][]int)
	for i := range nums {
		nmap[nums[i]] = append(nmap[nums[i]], i)
		res := nmap[nums[i]]
		if len(res) >= 2 {
			for j := 1; j < len(res); j++ {
				if res[j]-res[j-1] <= k {
					return true
				}
			}
		}
	}
	return false
}
