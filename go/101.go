package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// recursion
func isSymmetric(root *TreeNode) bool {
	return isSameTree(root, root)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
	}
}

// iteration
func isSymmetric(root *TreeNode) bool {
	n1 := root
	n2 := root
	queue := []*TreeNode{}
	queue = append(queue, n1)
	queue = append(queue, n2)
	for len(queue) > 0 {
		n1, n2 := queue[0], queue[1]
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)
		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
	}
	return true
}
