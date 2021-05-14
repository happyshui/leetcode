package leetcode

func subsets(nums []int) [][]int {
	var result [][]int
	var res []int
	backtrack(nums, 0, res, &result)
	return result
}

func backtrack(nums []int, idx int, res []int, result *[][]int) {
	ans := make([]int, len(res))
	copy(ans, res)
	*result = append(*result, ans)
	for i := idx; i < len(nums); i++ {
		res = append(res, nums[i])
		backtrack(nums, i+1, res, result)
		res = res[0 : len(res)-1]
	}
}
