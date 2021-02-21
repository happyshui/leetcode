package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	var length int
	for cur != nil {
		length++
		cur = cur.Next
	}
	virtual := &ListNode{0, head}
	newcur := virtual
	for i := 0; i < length-n; i++ {
		newcur = newcur.Next
	}
	newcur.Next = newcur.Next.Next
	return virtual.Next
}

// slow fast
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	virtual := &ListNode{0, head}
	fast, slow := head, virtual
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return virtual.Next
}
