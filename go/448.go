package leetcode

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	nmap := make(map[int]bool)
	for i := range nums {
		nmap[nums[i]] = true
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		_, ok := nmap[i+1]
		if !ok {
			result = append(result, i+1)
		}
	}
	return result
}
