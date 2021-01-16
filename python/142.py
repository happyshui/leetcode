#!/usr/bin/python3

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        slow = fast = head
        addr = []
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                cur = head
                while cur:
                    addr.append(cur)
                    cur = cur.next
                    if cur in addr:
                        return cur
        return None
