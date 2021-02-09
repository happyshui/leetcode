package leetcode

func summaryRanges(nums []int) []string {
	result := []string{}
	length := len(nums)
	if length == 0 {
		return result
	}
	if length == 1 {
		result = append(result, strconv.Itoa(nums[0]))
		return result
	}
	for i := 0; i < length; {
		start := i
		for i = i + 1; i < length && nums[i]-nums[i-1] == 1; i++ {
		}

		str := strconv.Itoa(nums[start])

		if start < i-1 {
			str = str + "->" + strconv.Itoa(nums[i-1])
		}
		result = append(result, str)
	}
	return result
}
