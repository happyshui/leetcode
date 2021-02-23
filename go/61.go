package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}
	cur := head
	tail := head
	var length int = 1
	for cur != nil && cur.Next != nil {
		cur = cur.Next
		length++
	}
	cur.Next = head

	for i := 1; i <= length-(k%length)-1; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil

	return head
}
