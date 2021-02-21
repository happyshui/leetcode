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
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	PreOrder(root, &result)
	return result
}

func PreOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if root != nil {
		*result = append(*result, root.Val)
		PreOrder(root.Left, result)
		PreOrder(root.Right, result)
	}
}

// interation
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}
