package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	depth(root, &result)
	return result - 1
}

func depth(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, result)
	right := depth(root.Right, result)
	*result = max(*result, left+right+1)
	return max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
