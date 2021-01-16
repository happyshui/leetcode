#!/usr/bin/python3

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        slow = fast = head
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                return True
        return False
