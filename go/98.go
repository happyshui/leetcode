package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	result := []int{}
	InfixOrder(root, &result)
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			return false
		}
	}
	return true
}

func InfixOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root != nil {
		InfixOrder(root.Left, result)
		*result = append(*result, root.Val)
		InfixOrder(root.Right, result)
	}
}
