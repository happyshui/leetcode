package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历k次
var result int

func kthSmallest(root *TreeNode, k int) int {
	InfixOrder(root, &k)
	return result
}

func InfixOrder(root *TreeNode, k *int) {
	if root != nil {
		InfixOrder(root.Left, k)
		*k--
		if *k == 0 {
			result = root.Val
		}
		InfixOrder(root.Right, k)
	}
}
