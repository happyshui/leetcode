package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}
	nodemap := make(map[*ListNode]int)
	cura := headA
	curb := headB
	for cura != nil {
		nodemap[cura] = 1
		cura = cura.Next
	}

	for curb != nil {
		_, ok := nodemap[curb]
		if ok {
			return curb
		}
		curb = curb.Next
	}
	return nil
}
