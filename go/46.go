package leetcode

func permute(nums []int) [][]int {
	var result [][]int
	var res []int
	visited := make([]bool, len(nums))
	backtrack(nums, visited, res, &result)
	return result
}

func backtrack(nums []int, visited []bool, res []int, result *[][]int) {
	if len(res) == len(nums) {
		ans := make([]int, len(res))
		copy(ans, res)
		*result = append(*result, ans)
		return
	} else {
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			res = append(res, nums[i])
			visited[i] = true
			backtrack(nums, visited, res, result)
			visited[i] = false
			res = res[0 : len(res)-1]
		}
	}
}
