package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}

// interation
func maxDepth(root *TreeNode) int {
	var count int
	if root == nil {
		return count
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count++
	}
	return count
}
