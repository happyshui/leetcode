package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// DFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

// BFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		n1, n2 := queue1[0], queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if n1.Val != n2.Val {
			return false
		}
		if (n1.Left == nil && n2.Left != nil) || (n1.Right == nil && n2.Right != nil) {
			return false
		}
		if (n1.Left != nil && n2.Left == nil) || (n1.Right != nil && n2.Right == nil) {
			return false
		}
		if n1.Left != nil {
			queue1 = append(queue1, n1.Left)
		}
		if n1.Right != nil {
			queue1 = append(queue1, n1.Right)
		}
		if n2.Left != nil {
			queue2 = append(queue2, n2.Left)
		}
		if n2.Right != nil {
			queue2 = append(queue2, n2.Right)
		}
	}
	return true
}
