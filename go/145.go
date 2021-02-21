package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postOrder(root, &result)
	return result
}

func postOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if root != nil {
		postOrder(root.Left, result)
		postOrder(root.Right, result)
		*result = append(*result, root.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}
	var visited *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == visited {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			visited = node
		} else {
			root = node.Right
		}
	}
	return result
}
