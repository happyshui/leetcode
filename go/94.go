package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	InfixOrder(root, &result)
	return result
}

// recursion
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

// iteration
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}

// divideAndConquer
func inorderTraversal(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	// Divide
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// Conquer
	result = append(result, left...)
	result = append(result, root.Val)
	result = append(result, right...)

	return result
}
