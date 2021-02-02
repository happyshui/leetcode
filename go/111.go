package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmin := minDepth(root.Left)
	rmin := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return lmin + rmin + 1
	} else {
		return min(lmin, rmin) + 1
	}
}

func min(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
