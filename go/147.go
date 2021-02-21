package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	virtual := &ListNode{Val: 0}
	pre := virtual
	cur := head
	for cur != nil {
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		next := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		pre = virtual
		cur = next
	}
	return virtual.Next
}
