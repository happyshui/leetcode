#!/usr/bin/python3

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        cur, prev = head, None
        while cur:
            # cur, cur.next, prev = cur.next, prev, cur
            cur.next, prev, cur = prev, cur, cur.next
        return prev
