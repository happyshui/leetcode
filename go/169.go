package leetcode

func majorityElement(nums []int) int {
	emap := make(map[int]int)
	var me int
	for i := 0; i < len(nums); i++ {
		_, ok := emap[nums[i]]
		if ok {
			emap[nums[i]] += 1
		} else {
			emap[nums[i]] = 1
		}
	}
	for k, v := range emap {
		if v > len(nums)/2 {
			me = k
		}
	}
	return me

}

// 进阶 时间复杂度O(n) 空间复杂度O(1)
